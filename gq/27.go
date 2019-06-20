package gq

import "fmt"

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}

func removeElement(nums []int, val int) int {
	n := 0
	for _, e := range nums {
		if e != val {
			nums[n] = e
			n++
		}
	}
	return n
}
