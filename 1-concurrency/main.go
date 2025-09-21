package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := make(chan int)
	squared := make(chan int)
	go generateItems(nums)
	go square(squared, nums)
	for n := range squared {
		fmt.Println(n)
	}
}

func generateItems(ch chan int) {
	for range 10 {
		ch <- rand.Intn(101)
	}
	close(ch)
}

func square(ch chan int, nums chan int) {
	for n := range nums {
		ch <- n * n
	}
	close(ch)
}
