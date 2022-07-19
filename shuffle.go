package main

import (
	"math/rand"
	"time"
)

//第二题
func shuffle() []int {
	arr := []int{0}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 54; i++ {
		arr[i] = i
	}
	for i := 53; i > 0; i-- {
		position := rand.Intn(i + 1)
		arr[position], arr[i] = arr[i], arr[position]
	}
	return arr
}
