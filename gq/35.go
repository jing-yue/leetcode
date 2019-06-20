package gq

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1,3}, 2))
}
func searchInsert(nums []int, target int) int {
	if len(nums) > 0 && target <= nums[0] {
		return 0
	}
	l, r := 0, len(nums)-1
	if l==r {
		if nums[l]>target {
			return l-1
		}else {
			return l+1
		}
	}
	for ; l < r; {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
			if l >= r {
				return l + 1
			}
		} else {
			r = mid - 1
			if l >= r {
				return r + 1
			}
		}
	}
	return 0
}
