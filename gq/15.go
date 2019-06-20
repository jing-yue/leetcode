package gq

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) (res [][]int) {
	l := len(nums)
	sort.Ints(nums)
	if nums[0] <= 0 && nums[l-2] >= 0 { //优化1: 整个数组同符号，则无解
		for i := 0; i < l-2; {
			if nums[i] > 0 {
				break // 最左值为正数则一定无解
			}
			first := i + 1
			last := l - 1
			for ; nums[i] != nums[i+1]; i++ {
				if first >= last || nums[i]*nums[last] > 0 {
					break // 两人选相遇，或者三人同符号，则退出
				}
				result := nums[i] + nums[first] + nums[last]
				if result == 0 {
					res = append(res, []int{nums[i], nums[first], nums[last]})
				}
				if result <= 0 { // 实力太弱，把菜鸟那边右移一位
					// 如果相等就跳过
					for ; first < last && nums[first] == nums[first+1]; first++ {

					}
				} else {
					// 实力太强，把大神那边右移一位
					for ; first < last && nums[last] == nums[last-1]; last-- {

					}
				}

				for ; nums[i] == nums[i+1]; i++ {

				}
			}
		}
	}
	return res
}
