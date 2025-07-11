package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	sorted := QuickSort(arr)
	wanted := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}

	if !SlicesEqual(sorted, wanted) {
		t.Errorf("Expected %v, got %v", wanted, sorted)
	}
}
