package main

import "fmt"

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		backup := arr[i]
		j := i - 1
		for j >= 0 && backup < arr[j] {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = backup
	}
}

func main() {
	arr := []int{1, 6, 7, 4, 2, 3, 10, 22, 5}
	insertSort(arr)
	fmt.Printf("arr: %+v", arr)
}
