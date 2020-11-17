package main

import "fmt"

// s = "aab"
func partition(s string) [][]string {
	result := make([][]string, 0)
	current := make([]string, 0)
	process(s, &result, current)
	return result
}

func process(s string, result *[][]string, current []string) {
	// fmt.Println(s)
	if len(s) == 0 {
		// fmt.Printf("save: %v\n", current)
		// *result = append(*result, current)
		// current = make([]string, 0)
		// return
		tempPath := make([]string, len(current))
		copy(tempPath, current)
		*result = append(*result, tempPath)
		return
	}
	for i := 1; i <= len(s); i++ {
		if !isPalindrome(s[:i]) {
			continue
		}
		current = append(current, s[:i])
		process(s[i:], result, current)
		current = current[:len(current)-1]
	}
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for j > i {
		if s[j] == s[i] {
			j--
			i++
		} else {
			return false
		}
	}
	return true
}
func main() {
	a := partition("aab")
	fmt.Print(a)
}
