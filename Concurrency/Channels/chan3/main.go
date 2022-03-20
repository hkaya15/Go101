package main

import (
	"fmt"
	"time"
)

type Deneme struct {
	a int
	b string
}

func scheduledNotification(t time.Duration) <-chan Deneme {
	ch := make(chan Deneme, 1)
	go func() {
		time.Sleep(t)
		ch <- Deneme{
			a: 3,
			b: "Struct",
		}
	}()
	return ch
}

func main() {
	fmt.Println("send first")
	a := <-scheduledNotification(time.Second)
	fmt.Println(a)
	fmt.Println("secondly send")
	b:=<-scheduledNotification(time.Second)
	fmt.Println(b)
	fmt.Println("lastly send")
}
