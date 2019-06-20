package gq

import "fmt"

func main() {
	fmt.Println(isPalindrome(123321))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := fmt.Sprintf("%d", x)
	s1, s2 := s[:len(s)/2], ""
	if len(s)%2 == 0 {
		s2 = s[len(s)/2:]
	} else {
		s2 = s[len(s)/2+1:]
	}

	for i, v := range s1 {
		if v != int32(s2[len(s2)-i-1]) {
			return false
		}
	}
	return true
}
