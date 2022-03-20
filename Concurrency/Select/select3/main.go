package main

import (
	"fmt"
	"time"
)

func main() {
	k1 := make(chan string)
	k2 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		k1 <- "video"
	}()
	go func() {
		time.Sleep(time.Second * 3)
		k2 <- "Ses"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m := <-k1:
			fmt.Println(m)
		case m1 := <-k2:
			fmt.Println(m1)
		}
	}
}
