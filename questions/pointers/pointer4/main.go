package main

import "fmt"

func main() {
	slice := make([]string, 1, 3)
	memoryfirstplace1 := &slice[0]
	fmt.Printf("İlk Adres: %v\n", memoryfirstplace1)
	func(slice []string) {
		slice = slice[1:3]
		slice[0] = "b"
		slice[1] = "b"
		fmt.Println(len(slice))
		fmt.Println(slice)
		memoryplace1 := &slice[0]
		memoryplace2 := &slice[1]
		fmt.Printf("1. Adres: %v", memoryplace1)
		fmt.Println()
		fmt.Printf("2. Adres: %v", memoryplace2)
	}(slice)
	fmt.Println()
	fmt.Println(len(slice))
	fmt.Println(slice)
	memoryplace1 := &slice[0]
	fmt.Printf("1. Adres: %v", memoryplace1)

}