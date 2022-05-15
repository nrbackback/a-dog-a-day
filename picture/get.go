package picture

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/nrbackback/a-dog-a-day/log"
)

const (
	layoutISO = "2006-01-02"
)

// Config to get picture
type Config struct {
	Keyword        string    `yaml:"keyword"`
	NotifyBeginDay string    `yaml:"notify_begin_day"`
	PictureFile    string    `yaml:"picture_file"`
	notifyBeginDay time.Time `yaml:"-"`
}

var pictureService struct {
	config Config
}

// TODO: 考虑挂载
func Init(c Config) error {
	notifyBeginDay, err := time.Parse(layoutISO, c.NotifyBeginDay)
	if err != nil {
		return err
	}
	pictureService.config = Config{
		Keyword:        c.Keyword,
		notifyBeginDay: notifyBeginDay,
		PictureFile:    c.PictureFile,
	}
	return nil
}

// CurrentPicture return which picture to push today
func CurrentPicture() (string, error) {
	now := time.Now()
	c := pictureService.config
	dayNow := now.Sub(c.notifyBeginDay).Hours() / 24
	dayNow++
	log.Logger.Infof("day now: %f", dayNow)
	picture, err := pictureByOffset(int(dayNow))
	if err != nil {
		return "", err
	}
	if err == nil && picture == "" {
		// fetch again
		picture, err = pictureByOffset(0)
		if err != nil {
			return "", err
		}
	}
	return picture, nil
}

func pictureByOffset(dayNow int) (string, error) {
	c := pictureService.config
	_, err := os.Stat(c.PictureFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := fetchPicture(dayNow); err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}
	f, err := os.Open(c.PictureFile)
	defer f.Close()
	if err != nil {
		return "", err
	}
	reader := bufio.NewReader(f)
	i := 0
	var picture string
	for {
		i++
		l, _, err := reader.ReadLine()
		if err != nil {
			if err != io.EOF {
				return "", err
			}
			if err := fetchPicture(i); err != nil {
				return "", err
			}
		}
		if i == int(dayNow) {
			picture = string(l)
			break
		}
	}
	return picture, nil
}

const requestURL = "https://image.baidu.com/search/acjson?tn=resultjson_com&ipn=rj&fp=result&word=%s&queryWord=%s&cl=2&lm=-1&ie=utf-8&oe=utf-8&st=-1&face=0&istype=2&nc=1&pn=%d&rn=%d"

// MaxPictureOnce get MaxPictureOnce picture in one request
const MaxPictureOnce = 50

func fetchPicture(nowLineCount int) error {
	c := pictureService.config
	pn := (nowLineCount / MaxPictureOnce) * MaxPictureOnce
	url := fmt.Sprintf(requestURL, c.Keyword, c.Keyword, pn, MaxPictureOnce)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var response struct {
		Data []struct {
			ThumbURL string `json:"thumbURL"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return err
	}

	f, err := os.OpenFile(c.PictureFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	contentToWrite := ""
	for i := 0; i < len(response.Data); i++ {
		if i == len(response.Data)-1 {
			contentToWrite += response.Data[i].ThumbURL
		} else {
			contentToWrite += response.Data[i].ThumbURL + "\n"
		}
	}
	if _, err := f.WriteString(contentToWrite); err != nil {
		return err
	}
	return nil
}
