package main

import "fmt"

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//示例:
//  给定 nums = [2, 7, 11, 15], target = 9
//  因为 nums[0] + nums[1] = 2 + 7 = 9
//  所以返回 [0, 1]

//solution01
func TwoSum(nums []int, target int) []int {
	var sub int
	aMap := make(map[int]int)
	for i, arrI := range nums {
		sub = target - arrI
		if _, ok := aMap[sub]; ok {
			return []int{aMap[sub], i}
		} else {
			aMap[arrI] = i
		}
	}
	return nil
}

//solution02
func TwoSum02(nums []int, target int) []int {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if target == v+nums[j] {
				return []int{i, j}
			}
		}
	}

	return nil
}

func main() {
	tests := [][]int{
		[]int{3, 2, 4}, //one
		[]int{1, 2, 3},
		[]int{1, 2},
		[]int{1, 5, 0, 15, 3, 2},
	}
	targets := []int{
		6,
		4,
		5,
		3,
	}
	results := [][]int{
		[]int{1, 2},
		[]int{0, 2},
		[]int{-1, -1},
		[]int{2, 4},
	}
	for i, _ := range tests {
		if ret := TwoSum(tests[i], targets[i]); ret != nil && ret[0] != results[i][0] && ret[1] != results[i][1] {
			fmt.Printf("TowSum01 case %d failes :%v\n", i, ret)
		}
		if ret := TwoSum02(tests[i], targets[i]); ret != nil && ret[0] != results[i][0] && ret[1] != results[i][1] {
			fmt.Printf("TowSum02 case %d failes :%v\n", i, ret)
		}
	}
}
