package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Println("Main execution started")
	fmt.Println("NO. of CPUs available:", runtime.NumCPU())
	fmt.Println("NO. of Goroutines available:", runtime.NumGoroutine())

	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.NumCPU())
}