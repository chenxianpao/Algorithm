package main

import "fmt"
import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	index := 0
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[index]
		index++
	}
	sort.Ints(nums1)
	fmt.Print(nums1)
}

func main() {
	input1 := []int{1, 2, 3, 0, 0, 0}
	input2 := []int{2, 5, 6}
	merge(input1, 3, input2, 3)
	// fmt.Print((a))
}
