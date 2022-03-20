package main

import (
	"fmt"
	"time"
)

func main() {
	// Create and read
	c := make(chan int, 10)
	go func() {
		defer close(c)
		time.Sleep(time.Second)
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
	// Range
	for val := range c {
		fmt.Println(val)
	}
	value, ok := <-c
	fmt.Println("Kanal'ın değeri: ", value, "Kapalılık", ok)
}
