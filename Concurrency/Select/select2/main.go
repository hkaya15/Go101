package main

import (
	"fmt"
	"time"
)

var c <-chan int
func main(){
	select{
	case <-c:
	case <-time.After(1*time.Second):
		fmt.Println("Timed out.")
	}
}