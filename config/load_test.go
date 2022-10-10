package config

import (
	"testing"
	"time"

	"github.com/nrbackback/a-dog-a-day/picture"
	"github.com/nrbackback/a-dog-a-day/runner"
	"github.com/nrbackback/a-dog-a-day/title"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(Load("../config.example.yml"))
	exceptConfig := GlobalConfig{
		Picture: picture.Config{
			Keyword:        "可爱狗狗",
			NotifyBeginDay: "2022-04-24",
			PictureFile:    "picture.txt",
		},
		Runner: runner.Config{
			NotifyTime:     "03:33PM",
			NotifyInterval: 24 * time.Hour,
		},
		Title: title.Config{
			Rand:      true,
			TitleFile: "title.txt",
		},
	}
	assert.Equal(exceptConfig, Config)
}
