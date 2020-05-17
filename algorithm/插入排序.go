package main

import (
	"fmt"
	"math/rand"
)

// 插入排序： 对于未排序的数据从前往后，每个数从后往前和已经排序好的数比较，如果比它前一个数小则交换位置。
func main() {
	//arr := []int{3, 4, 8, 9, 1, 7, 6, 5, 2, 0}
	n := 1000
	arr := generateArr(n)
	sortedArr := insertSort(arr, n)
	fmt.Println(sortedArr)
}

func generateArr(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		tmp := rand.Intn(n)
		arr[i] = tmp
	}
	return arr
}

func insertSort(arr []int, n int) (sortedArr []int) {
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}
