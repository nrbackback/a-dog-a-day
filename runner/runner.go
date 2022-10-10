package runner

import (
	"time"

	"github.com/nrbackback/a-dog-a-day/email"
	"github.com/nrbackback/a-dog-a-day/log"
	"github.com/nrbackback/a-dog-a-day/picture"
	"github.com/nrbackback/a-dog-a-day/title"
)

const DayLayout = "2006-01-02"

const TimeLayout = "03:00:30"

type Config struct {
	NotifyTime     string        `yaml:"notify_time"`
	NotifyInterval time.Duration `yaml:"notify_interval"`
}

var exitChanel = make(chan bool)

func Exit() {
	exitChanel <- true
}

func Start(c Config) {
	title, err := title.RandTitle()
	if err != nil {
		log.Logger.Error(err)
		return
	}
	pictureURL, err := picture.CurrentPicture()
	if err != nil {
		log.Logger.Error(err)
	} else {
		email.Send(title, pictureURL)
		log.Logger.Info("notify pushed")
	}
}
