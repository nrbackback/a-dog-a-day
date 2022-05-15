package title

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRandTitle(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	c := Config{
		Rand: false,
	}
	Init(c)
	title, err := RandTitle()
	require.NoError(err)
	assert.Equal("今日推送", title)
	c = Config{
		Rand:      true,
		TitleFile: "../title.example.txt",
	}
	Init(c)
	title, err = RandTitle()
	require.NoError(err)
	t.Log(title)
}
