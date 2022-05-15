package runner

import (
	"sync"
	"time"

	"github.com/nrbackback/a-dog-a-day/log"
	"github.com/nrbackback/a-dog-a-day/picture"
	"github.com/nrbackback/a-dog-a-day/push"
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

func Start(c Config, wg *sync.WaitGroup) {
	defer wg.Done()
	notifyTime, err := time.Parse(time.Kitchen, c.NotifyTime)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	title, err := title.RandTitle()
	if err != nil {
		log.Logger.Error(err)
		return
	}
	now := time.Now()
	firstNotifyTime := time.Date(now.Year(), now.Month(), now.Day(),
		notifyTime.Hour(), notifyTime.Minute(), notifyTime.Second(), 0, time.Local)
	if firstNotifyTime.Before(now) {
		firstNotifyTime = firstNotifyTime.Add(c.NotifyInterval)
	}
	log.Logger.Infof("notify will be sent at %s", firstNotifyTime.Format(time.RFC3339))
	select {
	case <-exitChanel:
		log.Logger.Info("runner exited normally")
		return
	case <-time.After(firstNotifyTime.Sub(now)):
		pictureURL, err := picture.CurrentPicture()
		if err != nil {
			log.Logger.Error(err)
		} else {
			push.Push(title, pictureURL)
			log.Logger.Info("notify pushed")
		}
	}

	ticker := time.NewTicker(c.NotifyInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			pictureURL, err := picture.CurrentPicture()
			if err != nil {
				log.Logger.Error(err)
			} else {
				log.Logger.Info("notify pushed")
				push.Push(title, pictureURL)
			}
		case <-exitChanel:
			log.Logger.Info("exit normally")
			return
		}
	}
}
