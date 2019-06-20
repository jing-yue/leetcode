package gq

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	maxarea := 0
	l := 0
	r := len(height) - 1

	for ; l < r; {
		h := 0
		sum := 0
		if height[l] < height[r] {
			h = height[l]
			sum = h * (r - l)
			l++
		} else {
			h = height[r]
			sum = h * (r - l)
			r--
		}
		if sum > maxarea {
			maxarea = sum
		}
	}

	return int(maxarea)
}

func maxArea1(height []int) int {
	if len(height) == 0 {
		return 0
	}
	res := 0
	for weight := len(height) - 1; weight > 0; weight-- {
		h := 0
		for x1 := 0; x1 < len(height)-weight; x1++ {
			x2 := weight + x1
			if height[x1] <= height[x2] && height[x1] > h {
				h = height[x1]
			}
			if height[x1] > height[x2] && height[x2] > h {
				h = height[x2]
			}
		}
		if weight*h > res {
			res = weight * h
		}
	}
	return res
}
