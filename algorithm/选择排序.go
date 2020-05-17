package main

import "fmt"

/*
选择排序： 从一组数据中找出最小或者最大的值与已经排序好的数据后的第一个数据交换位置。
*/

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	n := 10
	sortedArr := selectSort(arr, n)
	fmt.Println(sortedArr)
}

func selectSort(arr []int, n int) (sortedArr []int) {
	for i := 0; i < n; i++ {
		// 寻找[i, n)之间的最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// 交换
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}
