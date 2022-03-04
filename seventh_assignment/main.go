package main

import "fmt"

// 1 type myType int

type rectangle struct{
	a,b int
}

// 4 type
// type family struct{
// 	name string
// 	age int
// 	isMarried bool
// }


func main(){
	// 1- Create type of int underlying type
	// var x myType=10

	// fmt.Printf("%T, %v\n",x,x)


	// 2- Create a x value which start value equals 10, after change the value through with pointer
	// x:=10
	// y:=&x
	// fmt.Println(y)
	// *y=22
	// fmt.Println(x)


	// 3- Create a Rectangle type with underlying struct and create display, area,circumference methods
	rec:=rectangle{a: 3, b:8}
	rec.display()
	fmt.Println(rec.area())
	fmt.Println(rec.circumference())


	// 4- Create Family Type

	//families:=showFamily()
	// for i:=0;i<len(families);i++{
	// 	fmt.Println(families[i])
	// }
	
	// for i,v := range families{
	// 	fmt.Println(i,v)
	// }

	
}

func(r rectangle) display(){
	fmt.Println(r.a , r.b)
}

func (r rectangle) area() int{
	return r.a * r.b
}

func (r rectangle) circumference() int{
	return 2*(r.a + r.b)
}

// func showFamily() []family{
	
// 	family1 := family{
// 		name: "HK",
// 		age: 29,
// 		isMarried: true,
// 	}

// 	family2 := family{
// 		name:"Gölge",
// 		age: 7,
// 	}

// 	var family3 family
// 	family3.age=40

// 	return []family{family1,family2,family3}
// }