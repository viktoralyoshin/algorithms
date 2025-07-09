package search

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 10, 11, 23}
	target := 11
	wantErr := false
	wantIdx := 6

	gotIdx, err := BinarySearch(arr, target)
	if (err != nil) != wantErr {
		t.Errorf("Binary Search returned error: %s", err)
		return
	}

	if wantIdx != gotIdx {
		t.Errorf("Binary Search = %d, wanted = %d", gotIdx, wantIdx)
	}
}
