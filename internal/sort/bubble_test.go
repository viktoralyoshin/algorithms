package sort

import (
	"testing"
)

func SlicesEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func TestBubbleSort(t *testing.T) {
	arr := []int{10, 3, 52, 41, 4}
	wantArr := []int{3, 4, 10, 41, 52}

	gotArr := BubbleSort(arr)

	if !SlicesEqual(gotArr, wantArr) {
		t.Errorf("BubbleSort = %v, wanted = %v", gotArr, wantArr)
	}
}
