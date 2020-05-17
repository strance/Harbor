package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Int()
	fmt.Println(n)

	arr := make([]int, 100)
	arr[0] = n
	fmt.Println(arr)

}
