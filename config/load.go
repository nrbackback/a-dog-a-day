package config

import (
	"io/ioutil"
	"time"

	"github.com/nrbackback/a-dog-a-day/picture"
	"github.com/nrbackback/a-dog-a-day/push"
	"github.com/nrbackback/a-dog-a-day/runner"
	"github.com/nrbackback/a-dog-a-day/title"
	"gopkg.in/yaml.v3"
)

// GlobalConfig global config
type GlobalConfig struct {
	Picture  picture.Config  `yaml:"picture"`
	Runner   runner.Config   `yaml:"runner"`
	Title    title.Config    `yaml:"title"`
	Push     push.PushConfig `yaml:"push"`
	NotifyAt time.Time       `yaml:"-"`
}

var Config GlobalConfig

func Load(file string) error {
	r, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(r, &Config)
	if err != nil {
		panic(err)
	}
	return nil
}
