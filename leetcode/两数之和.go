package main

import "fmt"

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

func main() {
	nums := []int{0, 1, 2, 3, 4}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}


func twoSum(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		func (i, tmp int) {
			for j := i; j < len(nums); j ++ {
				if nums[j] == tmp && nums[j] != nums[i]{
					res = []int{i, j}
					break
				}
			}
		}(i, tmp)
	}

	return res
}