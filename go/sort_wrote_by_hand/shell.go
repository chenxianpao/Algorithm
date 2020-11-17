package main

import "fmt"

func shellSort(arr []int, start, gap int) {
	for i := start + gap; i < len(arr); i += gap {
		backup := arr[i]
		j := i - gap
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j]
			j -= gap
		}
		arr[j+gap] = backup
	}
}

func sort(arr []int) {
	gap := len(arr) / 2
	for gap > 0 {
		for pos := 0; pos < gap; pos++ {
			shellSort(arr, pos, gap)
		}
		gap /= 2
	}
}
func main() {
	arr := []int{1, 6, 7, 4, 2, 3, 10, 22, 5}
	sort(arr)
	fmt.Printf("arr: %+v", arr)
}
