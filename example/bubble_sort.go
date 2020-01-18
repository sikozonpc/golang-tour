package main

import "fmt"

func main() {
	arr := []int{5, 1, 4, 2, 8}

	fmt.Println(BubbleSort(arr))
}

// BubbleSort : Algorithm to sort an array of arr with O(n^2) complexity
func BubbleSort(arr []int) []int {
	for i := len(arr); i > 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j-1] > arr[j] {
				intermediate := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = intermediate
			}
		}
	}
	return arr
}
