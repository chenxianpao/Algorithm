package main

import "fmt"

func bubble(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
}

func main() {
	arr := []int{1, 6, 7, 4, 2, 3, 10, 22, 5}
	bubble(arr)
	fmt.Printf("arr: %+v", arr)
}
