package title

import (
	"bufio"
	"io"
	"math/rand"
	"os"
	"time"
)

type Config struct {
	Rand      bool   `yaml:"rand"`
	TitleFile string `yaml:"title_file"`
}

var titleService struct {
	config Config
}

func Init(c Config) {
	rand.Seed(time.Now().UnixNano())
	titleService.config = c
}
func RandTitle() (string, error) {
	c := titleService.config
	if !c.Rand {
		return "今日推送", nil
	}
	f, err := os.Open(c.TitleFile)
	if err != nil {
		return "", err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	allLines := make([]string, 0)
	count := 0
	for {
		l, _, err := r.ReadLine()
		if len(l) != 0 {
			count++
			allLines = append(allLines, string(l))
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	rand.Seed(time.Now().UnixNano())
	order := rand.Intn(count)
	return allLines[order], nil
}
