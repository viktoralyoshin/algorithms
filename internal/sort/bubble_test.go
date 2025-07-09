package sort

import (
	"testing"
)

func SlicesEqual(slc1, slc2 []int) bool {
	if len(slc1) != len(slc2) {
		return false
	}

	for i := range slc1 {
		if slc1[i] != slc2[i] {
			return false
		}
	}

	return true
}

func TestBubbleSort(t *testing.T) {
	slc := []int{10, 3, 52, 41, 4}
	wantSlc := []int{3, 4, 10, 41, 52}

	gotSlc := BubbleSort(slc)

	if !SlicesEqual(gotSlc, wantSlc) {
		t.Errorf("BubbleSort = %v, wanted = %v", gotSlc, wantSlc)
	}
}
