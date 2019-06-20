package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}

func singleNumber1(nums []int) int {
	hash := make(map[int]interface{})
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			delete(hash, nums[i])
		} else {
			hash[nums[i]] = struct{}{}
		}
	}

	for e := range hash {
		return e
	}
	return -1
}

func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i = i + 2 {
		if i+1 >= len(nums) {
			return nums[i]
		}
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return -1
}
