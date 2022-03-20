package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()
	check(ctx)
	fmt.Println("Final Write")
}

func check(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 2)
	for {
		select {
		case <-ticker.C:
			writeMessage()
		case <-ctx.Done():
			fmt.Println("Done")
			return
		}
	}
}

func writeMessage() {
	log.Println("this is message")
}
