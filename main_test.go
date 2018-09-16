package main

import (
	"reflect"
	"testing"
)

func TestCase(t *testing.T) {
	arr := []int{9, 6, 7, 8, 5, 4, 3, 2, 1}
	sortedArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(quickSort(arr), sortedArr) {
		t.Errorf("Expected: %v, Output: %v\n", sortedArr, quickSort(arr))
	}
}
