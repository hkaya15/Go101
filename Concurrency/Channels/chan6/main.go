package main

import (
	"fmt"
	"sync"
)

func main() {

	channel := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-channel
			fmt.Printf("%v başladı\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines")
	fmt.Println("Other jobs...")

	chanOwner := func() <-chan int {
		channel1 := make(chan int, 5)
		go func() {
			defer close(channel1)
			for i := 0; i < 5; i++ {
				channel1 <- i
			}
		}()
		return channel1
	}

	close(channel)
	res := chanOwner()
	for result := range res {
		fmt.Printf("Value: %v\n", result)
	}
	fmt.Println("Done")
	wg.Wait()

}
