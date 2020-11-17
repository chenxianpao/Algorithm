package main

import "fmt"

func majorityElement(nums []int) int {
	countMap := map[int]int{}
	for _, v := range nums {
		if _, ok := countMap[v]; !ok {
			countMap[v] = 1
		} else {
			countMap[v] += 1
		}
	}
	for k, v := range countMap {
		if v > len(nums)/2 {
			return k
		}
	}
	return -1
}

func main() {
	input := []int{3, 2, 3}
	a := majorityElement(input)
	fmt.Print((a))
}
