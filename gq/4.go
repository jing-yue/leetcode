package gq

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums:= append(nums1, nums2...)
	sort.Ints(nums)
	return mid(nums)
}

func mid(nums []int) float64 {
	l := len(nums)
	if l%2 == 0 {
		return float64(nums[l/2]+nums[l/2-1]) / 2.0
	} else {
		return float64(nums[l/2])
	}
}
