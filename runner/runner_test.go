package runner

import (
	"sync"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	c := Config{
		NotifyTime:     "05:16PM",
		NotifyInterval: 24 * time.Hour,
	}
	go Start(c, &wg)
	Exit()
	wg.Wait()
}
