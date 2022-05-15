package main

import (
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/nrbackback/a-dog-a-day/config"
	"github.com/nrbackback/a-dog-a-day/log"
	"github.com/nrbackback/a-dog-a-day/picture"
	"github.com/nrbackback/a-dog-a-day/push"
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
	push.Init(config.Config.Push)
	title.Init(config.Config.Title)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	var wg sync.WaitGroup
	wg.Add(2)
	log.Logger.Info("dog start...")
	go runner.Start(config.Config.Runner, &wg)
	go func() {
		<-sigs
		log.Logger.Info("interupted...")
		runner.Exit()
		wg.Done()
	}()
	wg.Wait()
}
