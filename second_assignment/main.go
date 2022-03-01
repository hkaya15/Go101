package main

import (
	"fmt"
	"strconv"
)



func main(){
	// 1- Type conversion
	// x:=75
	// var y float64
	// y=float64(x)
	// fmt.Printf("%T, %v\n",y,y)


	// 2- Multiple assing sample
	// x:=5
	// y:=10
	// fmt.Println(x,y)
	// x,y = y,x
	// fmt.Println(x,y)


	// 3- Non English variable names
	// yaş:=40
	// fmt.Println(yaş)


	// 4- Shadowing
	// x:=5
	// if true{
	// 	x=10
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)


	// 5- 40 as a string

	 x:=40
	// y:=string(x)
	// fmt.Println(y)

	s:= strconv.Itoa(x)
	fmt.Printf("%v,%T\n",s,s)

}