package main

import "fmt"

func singleNumber(nums []int) int {
	countMap := map[int]int{}
	for _, v := range nums {
		if _, ok := countMap[v]; !ok {
			countMap[v] = 1
		} else {
			countMap[v] += 1
		}
	}
	fmt.Print(countMap)
	for k, v := range countMap {
		if v == 1 {
			return k
		}
	}
	return -1
}

func singleNumber1(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {
	input := []int{2, 2, 1}
	a := singleNumber1(input)
	fmt.Print((a))
}
