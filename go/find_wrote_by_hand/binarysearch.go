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

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	benchmark := arr[start]
	i := start
	j := end
	for i < j {
		for arr[i] < benchmark && i < j {
			i++
		}
		for arr[j] > benchmark && i < j {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	quickSort(arr, start, i-1)
	quickSort(arr, i+1, end)
}

func main() {
	arr := []int{1, 6, 7, 4, 2, 3, 10, 22, 5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Printf("arr: %+v", arr)
	result := binarySearch1(arr, 1)
	fmt.Printf("result: %+v", result)
}
