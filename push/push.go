package push

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/nrbackback/a-dog-a-day/log"
)

var pushUrlFormat = "https://sctapi.ftqq.com/%s.send"

const (
	TitleParams = "title"
	DespParams  = "desp"
)

type PushConfig struct {
	ServerKey string `yaml:"server_key"`
}

var pushService struct {
	config PushConfig
}

func Init(c PushConfig) {
	pushService.config = c
}

// Push push once
func Push(title, image string) error {
	pushURL := fmt.Sprintf(pushUrlFormat, pushService.config.ServerKey)
	htpRequest, err := http.NewRequest(http.MethodGet, pushURL, nil)
	if err != nil {
		return err
	}
	q := htpRequest.URL.Query()
	q.Add(TitleParams, title)
	q.Add(DespParams, fmt.Sprintf("![url](%s)", image))
	htpRequest.URL.RawQuery = q.Encode()
	var client http.Client
	resp, err := client.Do(htpRequest)
	if err != nil {
		log.Logger.Error(err)
		return err
	}
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		msg := fmt.Sprintf("got unexcepted status code %d, body: %s", resp.StatusCode, string(body))
		return errors.New(msg)
	}
	return err
}
