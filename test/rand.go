package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i))
		order := rand.Intn(10)
		fmt.Println("-------order-----", order)
	}
}
