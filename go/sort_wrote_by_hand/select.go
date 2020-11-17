package main

import "fmt"

func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j:= i+1; j<len(arr);j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	arr := []int{1, 6, 7, 4, 2, 3, 10, 22, 5}
	selectSort(arr)
	fmt.Printf("arr: %+v", arr)
}
