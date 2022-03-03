package main

import "fmt"

func main() {
	slice := make([]string, 2, 3)
	memoryfirstplace1 := &slice[0]
	memorysecondplace := &slice[1]
	fmt.Printf("İlk Adres: %v", memoryfirstplace1)
	fmt.Println()
	fmt.Printf("İkinci Adres: %v", memorysecondplace)
	fmt.Println()
	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Println()
		fmt.Println(slice)
		memoryplace1 := &slice[0]
		memoryplace2 := &slice[1]
		fmt.Printf("1. Adres: %v", memoryplace1)
		fmt.Println()
		fmt.Printf("2. Adres: %v", memoryplace2)
	}(slice)
	fmt.Println()
	fmt.Println()
	fmt.Println(slice)
	memoryplace1 := &slice[0]
	memoryplace2 := &slice[1]
	fmt.Printf("1. Adres: %v", memoryplace1)
	fmt.Println()
	fmt.Printf("2. Adres: %v", memoryplace2)
}
