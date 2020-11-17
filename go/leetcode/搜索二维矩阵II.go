package main

import "fmt"

func binarySearch1(arr []int, findValue int) int {
	left := 0
	right := len(arr) - 1
	for {
		if left > right {
			break
		}
		mid := (left + right) / 2
		if arr[mid] < findValue {
			left = mid + 1
		} else if arr[mid] > findValue {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func searchMatrix(matrix [][]int, target int) bool {
	// if len(matrix) == 1 && len(matrix[0]) == 1 {
	// 	if matrix[0][0] == target {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// }
	for i:=0; i< len(matrix); i++ {
		result := binarySearch(matrix[i], target)
		if result != -1 {
			return true
		}
	}
	return false
}

func main() {
	input := []int{3, 2, 3}
	a := searchMatrix(input)
	fmt.Print((a))
}
