package gq

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr), arr)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i]=nums[j]
		}
	}
	return i+1
}

func inArray(nums []int, tmp int) bool {
	for _, e := range nums {
		if e == tmp {
			return true
		}
	}
	return false
}

func removeDuplicates1(nums []int) int {
	k := 0
	for _, v := range nums {
		if !inArray(nums[:k], v) {
			nums[k] = v
			k++
		}
	}
	nums = nums[:k]
	return len(nums)
}
