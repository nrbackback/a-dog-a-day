package push

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestPush(t *testing.T) {
	require := require.New(t)
	title := "hi"
	serverKey := "test"
	image := "http://pic.bizhi360.com/bpic/43/5943.jpg"
	pushURL := fmt.Sprintf(pushUrlFormat, serverKey)

	defer gock.Off()
	gock.New(pushURL).
		Reply(200).
		JSON(map[string]string{"foo": "bar"})

	Init(PushConfig{serverKey})
	require.NoError(Push(title, image))
}
