package main

import (
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go time.Sleep(1e9)
	}
	time.Sleep(2e9)
}
