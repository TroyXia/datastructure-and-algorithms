package main

import "fmt"

func partition(arr []int, left, right int) int {
	pivot := arr[left]
	isRight := true
	for left < right {
		if isRight {
			if arr[right] > pivot {
				right--
			} else {
				arr[left] = arr[right]
				left++
				isRight = false
			}
		} else {
			if arr[left] > pivot {
				arr[right] = arr[left]
				right--
				isRight = true
			} else {
				left++
			}
		}
	}

	if left == right {
		arr[left] = pivot
	}

	return left
}

func QuickSort(arr []int, left, right int) {
	if left < right {
		k := partition(arr, left, right)

		QuickSort(arr, left, k-1)
		QuickSort(arr, k+1, right)
	}
}

func main() {
	var arr = []int{11, 8, 3, 9, 7, 1, 2, 5}
	QuickSort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}
