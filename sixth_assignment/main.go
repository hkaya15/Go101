package main

import (
	"fmt"
	"merhaba"
)

func main() {
	// 1- Create 5 string value array and 5 int value slice
	// myArr:=[5]string{"S.F","L.A","Munich","Dublin","Kosova"}

	// for i,v := range myArr{
	// 	fmt.Println(i,v)
	// }

	// mySlc:=[]int{1,2,3,4,5}
	// for i,v:= range mySlc{
	// 	fmt.Println(i,v)
	// }

	// 2- Slice slicing
	// mySlc:=[]int{0,1,2,3,4,5,6,7,8}
	// mySlc1:=mySlc[:4]
	// fmt.Println(mySlc1)

	// mySlc2:=mySlc[4:7]
	// fmt.Println(mySlc2)

	// mySlc3:=mySlc[6:]
	// fmt.Println(mySlc3)

	// mySlc4:=append(mySlc[2:4],mySlc[6:8]...)
	// fmt.Println(mySlc4)

	// 3- What's the difference copy method and assign for slices
	// mySlc:=[]int{1,2,3}
	// mySlc2:=make([]int,2)

	// //copy(mySlc2,mySlc)
	// mySlc2=mySlc // While assigning has a reference value, copy has a value
	// mySlc[0]=44
	// fmt.Println(mySlc)
	// fmt.Println(mySlc2)

	// 4- Map shows with for range
	myMap := map[string][]string{
		"Yaşar Kemal":    {"İnce Memed", "Yer Demir Gök Bakır"},
		"Sabahattin Ali": {"Kuyucaklı Yusuf", "Kürk Mantolu Madonna", "Değirmen"},
	}

	for key, value := range myMap {
		fmt.Println("Yazar: ", key)
		fmt.Println("Eserleri :")
		for _, v := range value {
			fmt.Println("\t", v)
		}
	}

	//Package Integration
	merhaba.Hello()
}
