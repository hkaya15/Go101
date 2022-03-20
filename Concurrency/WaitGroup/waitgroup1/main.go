package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	go f1(&wg)
	fmt.Println("İşlem Başlıyor!")
	wg.Wait()

}

func f1(wg *sync.WaitGroup){
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			fmt.Println(i)
		}(i)
		
	}

}