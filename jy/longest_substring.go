/*
	给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
	示例 1:

	输入: "abcabcbb"
	输出: 3
	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 */
package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("asdasdasdasdasdasdasdasdawqweqweqweqweqfdsfdsf"))
}

func lengthOfLongestSubstring(s string) int {
	val := []byte(s)
	kvMap := make([]int, 128)
	lens := len(s)
	var max, num int
	for i, j := 0, 0; i < lens && j < lens; j++ {
		if kvMap[val[j]] > i {
			i = kvMap[val[j]]
		}
		num = j - i + 1
		if num > max {
			max = num
		}
		kvMap[val[j]] = j + 1
	}
	return max
}