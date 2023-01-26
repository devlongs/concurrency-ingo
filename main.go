package main

import (
	"fmt"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup){
	fmt.Println("f1 goroutine execution started")
	for i := 0; i < 3; i++{
		fmt.Println("f1, i=", i)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func f2(){
	fmt.Println("f2 goroutine execution started")
	for i := 5; i < 8; i++{
		fmt.Println("f2, i=", i)
	}
}
func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	
	go f1(&wg)
	

	f2()

	fmt.Println("Main execution stopped")
}