package gq

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValid("(())"))
}

func isValid(s string) bool {
	list := []string{}

	hash := make(map[string]string)
	hash[")"] = "("
	hash["}"] = "{"
	hash["]"] = "["
	for _, e := range strings.Split(s, "") {
		if e == "(" || e == "{" || e == "[" {
			list = append(list, e)
		} else {
			if len(list) == 0 || list[len(list)-1] != hash[e] {
				return false
			} else {
				list = list[:len(list)-1]
			}
		}
	}
	return len(list) == 0
}
