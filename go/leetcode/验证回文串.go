package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	// s = strings.Replace(s, " ", "", -1)
	// test := []string{}
	// for _, v := range s {
	// 	if unicode.IsLetter(v) {
	// 		test = append(test, strings.ToLower(string(v)))
	// 	}
	// }
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1
	for left < right {
		for left < right && !isalnum((s[left])) {
			left++
		}
		for left < right && !isalnum((s[right])) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	// fmt.Print(string(s[0]))
	return true
}
func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func main() {
	// s := "A man, a plan, a canal: Panama"
	// s := "race a car"
	s := "OP"
	a := isPalindrome(s)
	fmt.Print((a))
}
