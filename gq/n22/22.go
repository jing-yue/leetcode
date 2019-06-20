package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}
func generateParenthesis(n int) (res []string) {
	if n == 0 {
		res = append(res, "")
		return res
	} else {
		for c := 0; c < n; c++ {
			for _, left := range generateParenthesis(c) {
				for _, right := range generateParenthesis(n - 1 - c) {
					res = append(res, "("+left+")"+right)
				}
			}
		}
	}
	return res
}
