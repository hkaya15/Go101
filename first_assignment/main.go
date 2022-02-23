package main

import "fmt"

// 1) Define variables with 3 different methods

// var studentName string = "John Doe"
// var grade = 77

// 2) Define variables in one line

// var (
// 	studentName string = "John Doe"
// 	grade       int    = 77
// 	isPassed           = true
// )

// var studentName, grade, isPassed = "John Doe",77,true

func main() {
	// 1)
	// isPassed := true

	//2)
	studentName,grade,isPassed := "John Doe", 77 , true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)
}
