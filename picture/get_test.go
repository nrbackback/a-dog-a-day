package picture

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}
	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)
		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func TestCurrentPicture(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	today := time.Now().Format(layoutISO)
	c := Config{
		Keyword:        "可爱狗狗",
		NotifyBeginDay: today,
		PictureFile:    "./picture.txt",
	}
	err := Init(c)
	require.NoError(err)
	p, err := CurrentPicture()
	require.NoError(err)
	assert.NotEmpty(p)
	f, err := os.Open(c.PictureFile)
	require.NoError(err)
	defer f.Close()
	count, err := lineCounter(f)
	require.NoError(err)
	assert.Equal(MaxPictureOnce, count)

	c.NotifyBeginDay = time.Now().AddDate(0, 0, -60).Format(layoutISO)
	err = Init(c)
	require.NoError(err)
	p, err = CurrentPicture()
	require.NoError(err)
	assert.NotEmpty(p)
	f, err = os.Open(c.PictureFile)
	require.NoError(err)
	defer f.Close()
	count, err = lineCounter(f)
	require.NoError(err)
	assert.Equal(2*MaxPictureOnce, count)
	require.NoError(os.Remove("./picture.txt"))

	err = Init(c)
	require.NoError(err)
	p, err = CurrentPicture()
	require.NoError(err)
	assert.NotEmpty(p)
	f, err = os.Open(c.PictureFile)
	require.NoError(err)
	defer f.Close()
	count, err = lineCounter(f)
	require.NoError(err)
	assert.Equal(2*MaxPictureOnce, count)
	require.NoError(os.Remove("./picture.txt"))
}
