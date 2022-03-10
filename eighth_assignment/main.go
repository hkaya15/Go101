package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type rectangle struct{
	a,b float64
}

type square struct{
	a float64
}

type shape interface{
	Area() float64
}

func (r rectangle) Area()float64{
	return r.a+r.b
}

func (s square) Area() float64{
	return s.a * s.a
}

func main(){
	rec:=rectangle{a: 4,b: 2}
	squ:=square{a: 3}
	shapes:= []shape{ rec, squ}
	for _,v := range shapes{
		fmt.Printf("%T tipli, %v alanı\n",v,v.Area())
	}


	// Type Interface

	// var i interface{}
	// // i=20
	// // i="hello"
	// i=struct{
	// 	FirstName string
	// 	LastName string
	// }{"Hüseyin","Kaya"}

	// fmt.Println(i)

	
	data := map[string]interface{}{}
	contents,err := ioutil.ReadFile("data.json")
	if err !=nil{
		panic(err)
	}
	if err:=json.Unmarshal(contents,&data); err!=nil{
		panic(err)
	}

	fmt.Println(data)
	fmt.Println(data["name"].(string))
	fmt.Println(data["age"].(float64))
	fmt.Println(data["details"].(map[string]interface{}))
}


type Person struct{
	Name string `json:"firstkey"`
	Age float64 `json:"age"`
	Details struct{
		FirstKey string `json:"firstKey"`
	} `json:"details"`
}