package main

import "fmt"

//Used for 4th question
//const x=24

func main(){
	// 1- x=x-10 && x-=10
	// x:=50
	// x-=10
	// fmt.Printf("%v, %T\n", x,x)

	// 2- Fahreneit to Kelvin
	// F:=-40
	// K:=float64(F-32)/1.8+273
	// fmt.Println(K)


	// 3-const
	// const myAge = 40
	// fmt.Printf("%T, %v\n",myAge,myAge)


	// 4- Are Const values work with "shadowing"?
	// const x=35
	// fmt.Println(x)
	// Answer is yes!


	// 5- const with different types(typless) sum
	// const x=4 // Typless
	// const y=5.4 // Typless

	// fmt.Println(x+y) // Untyped const values work with type conversion


	// 6- const with typed values sum
	const x float64=6.4
	y:=4+x
	fmt.Println(y)



}