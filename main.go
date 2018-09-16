package main

import "fmt"

func main() {
	arr := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	fmt.Println(quickSort(arr))
}

func quickSort(array []int) []int {
	quickSortHelper(array, 0, len(array)-1)
	return array
}

func quickSortHelper(array []int, startIndex, endIndex int) []int {
	if startIndex >= endIndex {
		return array
	}
	pivotIndex := startIndex
	leftIndex := startIndex + 1
	rightIndex := endIndex
	for rightIndex >= leftIndex {
		if array[leftIndex] > array[pivotIndex] && array[rightIndex] < array[pivotIndex] {
			swap(leftIndex, rightIndex, array)
		}
		if array[leftIndex] <= array[pivotIndex] {
			leftIndex++
		}
		if array[rightIndex] >= array[pivotIndex] {
			rightIndex--
		}
	}
	swap(pivotIndex, rightIndex, array)

	if rightIndex-1-startIndex < endIndex-(rightIndex-1) {
		quickSortHelper(array, startIndex, rightIndex-1)
		quickSortHelper(array, rightIndex+1, endIndex)
	} else {
		quickSortHelper(array, rightIndex+1, endIndex)
		quickSortHelper(array, startIndex, rightIndex-1)
	}
	return array
}

func swap(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}
