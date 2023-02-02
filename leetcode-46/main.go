package main

import "fmt"

func main() {
	fmt.Println(permute([]int{5, 4, 6, 2}))
}

var result [][]int

func permute(nums []int) [][]int {
	lens := len(nums)
	if lens == 1 {
		result = append(result, nums)
		return result
	}
	result = make([][]int, 0)
	res := make([]int, 0)
	backtrack(nums, lens, 0, res)
	return result
}

func backtrack(nums []int, lens int, index int, res []int) {
	if index == lens {
		fmt.Println("rrrrrrrrrrrrrrrr ", index, " --- ", res)
		r := make([]int, 0)
		r = append(r, res...)
		result = append(result, r)
		fmt.Println("result ", result)
		return
	}

	for i := 0; i < lens; i++ {
		if contain(res, nums[i]) {
			continue
		}
		res = append(res, nums[i])
		fmt.Println(index, " --- ", i, " ``````` ", res, " --- ", nums[i])
		backtrack(nums, lens, index+1, res)
		res = res[:len(res)-1]
	}
}

func contain(res []int, num int) bool {
	for _, v := range res {
		if v == num {
			return true
		}
	}
	return false
}
