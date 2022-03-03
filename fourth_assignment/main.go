package main

import "fmt"

func main(){
	// 1- print values from 1 to 10 with "if" statement regarding of the even or odd
	// for i:=1;i<11;i++{
	// 	if i%2==0{ 
	// 		fmt.Println(i,"çifttir")
	// 	}else{
	// 		fmt.Println(i,"tektir")
	// 	}
	// }


	// 2- use "for" statement to work as "while" statement that is not in GoLang 
	// x:=0
	// for x<10{
	// 	fmt.Println(x)
	// 	x++
	// }


	// 3- Switch fallthrough
	// switch x:=25;{
	// case x<20:
	// 	fmt.Printf("%d küçüktür 20'den\n",x)
	// 	fallthrough
	// case x<50:
	// 	fmt.Printf("%d küçüktür 50'den\n",x)
	// 	fallthrough
	// case x<100:
	// 	fmt.Printf("%d küçüktür 100'den\n",x)
	// 	fallthrough
	// case x<200:
	// 	fmt.Printf("%d küçüktür 200'den\n",x)
	// }


	// 4- Do idiomatic the "if" statement
	// if x:=20;x%2==0{
	// 	fmt.Println(x,"çifttir")
	// }else{
	// 	fmt.Println(x,"tektir")
	// }

	// x:=20
	// if x%2==0{
	// 	fmt.Println(x,"çifttir")
	// 	return
	// }
	// fmt.Println(x,"tektir")


	// 5- write a code that returns beetween 1 and 50 prime numbers
	var x,y int
	for x=2;x<50;x++{
		for y=2;y<50;y++{
			if x%y==0 {
				break
			}
		}
		if y>(x/y){
			fmt.Printf("%d bir asal sayıdır\n",x)
		}
	}
}