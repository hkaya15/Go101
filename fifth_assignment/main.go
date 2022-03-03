package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 1- write a funtion that return addition, substraction, multiplication beetween two values
	//fmt.Println(calculation(5,2))

	// 2- Return "pass" or "fail" value regarding of the user number input
	// fmt.Print("Lütfen aldığınız notu giriniz: ")
	// grade, _ :=getGrade()

	// var result string

	// if grade >= 50{
	// 	result="Geçtin"
	// }else{
	// 	result="Kaldın"
	// }
	// fmt.Println(result)

	// 3- Write a guessing code beetween 1 and 100 number. Total quess val=10
	
	target := numRand()
	fmt.Println("1 ile 100 arasındaki sayıyı bulmaya çalışın")
	reader := bufio.NewReader(os.Stdin)
	
	for attempts := 0; attempts < 10; attempts++ {
		fmt.Println(10-attempts, "hakkınız kaldı")
		fmt.Print("Lütfen tahmininizi yapınız: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}
		if attempts != 9 {
			if num > target {
				fmt.Println("Daha küçük bir sayı gir")
			} else if num < target {
				fmt.Println("Daha büyük bir sayı gir")
			} else {
				fmt.Println("Doğru Tahmin!!!", attempts+1, "seferde buldunuz.")
				break
			}
		} else {
			fmt.Println("Şans senden yana değil!!!")
		}

	}

}

// 1 ****
// func calculation(num1,num2 int)(sum,dif,multi int){
// 	sum=num1+num2
// 	dif=num1-num2
// 	multi=num1*num2

// 	return sum,dif,multi
// }

// 2 ****
// func getGrade()(int,error){
// 	reader:=bufio.NewReader(os.Stdin)
// 	input,err:=reader.ReadString('\n')
// 	if err!=nil{
// 		fmt.Println(err)
// 	}

// 	input=strings.TrimSpace(input)
// 	num,err:=strconv.Atoi(input)
// 	if err!=nil{
// 		fmt.Println(err)
// 	}
// 	return num,nil
// }

func numRand() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(100)
}
