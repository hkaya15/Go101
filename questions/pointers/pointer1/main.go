package main

import "fmt"

func main() {
	var k *int
	fmt.Printf("%v\n", *&k)
	fmt.Printf("%v\n", &k)
}