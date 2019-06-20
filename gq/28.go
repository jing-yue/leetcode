package gq

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strStr("aa", "aaa"))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}
	haystacks := strings.Split(haystack, "")
	needles := strings.Split(needle, "")
	l := len(needles)
	for i, e := range haystacks {
		if e == needles[0] {
			if i+l > len(haystack) {
				return -1
			}
			if needle == strings.Join(haystacks[i:i+l], "") {
				return i
			}
		}
	}
	return -1
}
