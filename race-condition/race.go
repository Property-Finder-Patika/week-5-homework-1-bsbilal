package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu sync.Mutex
)

func main() {
	numbers := []int{}
	/*
			This code gives an error for go race condition..
		go appendTwo(&numbers)
		go appendOne(&numbers)

	*/

	//this version of upper (without go routine('s)) code works fine
	appendOne(&numbers)
	appendTwo(&numbers)

	time.Sleep(1 * time.Millisecond)
}

func appendTwo(numbers *[]int) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("numbers %v\n", numbers)
	*numbers = append(*numbers, []int{1, 2}...)
	fmt.Printf("numbers second iter %v\n", numbers)
}

func appendOne(numbers *[]int) {
	mu.Lock()
	defer mu.Unlock()
	*numbers = append(*numbers, 12)
}
