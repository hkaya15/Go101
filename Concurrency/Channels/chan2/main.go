package main

import (
	"fmt"
	"time"
)

type T = struct{}

func main() {
	completed := make(chan T)
	go func() {
		fmt.Println("ping")
		time.Sleep(time.Second * 5) // heavy process simulation
		<-completed                 // receive a value from completed channel
	}()

	completed <- struct{}{} // blocked waiting for a notification
	fmt.Println("pong")
}
