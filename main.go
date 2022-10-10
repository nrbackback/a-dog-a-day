package main

import (
	"flag"

	"github.com/nrbackback/a-dog-a-day/config"
	"github.com/nrbackback/a-dog-a-day/email"
	"github.com/nrbackback/a-dog-a-day/log"
	"github.com/nrbackback/a-dog-a-day/picture"
	"github.com/nrbackback/a-dog-a-day/runner"
	"github.com/nrbackback/a-dog-a-day/title"
)

func main() {
	log.Init()
	var configFile string
	flag.StringVar(&configFile, "config", "config.yml", "load config")
	flag.Parse()

	if err := config.Load(configFile); err != nil {
		panic(err)
	}
	if err := picture.Init(config.Config.Picture); err != nil {
		panic(err)
	}
	email.WeiboSender = email.EmailSender{
		Conf: config.Config.Email,
	}
	title.Init(config.Config.Title)
	runner.Start(config.Config.Runner)
}
