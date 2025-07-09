package sort

import "testing"

func TestSelectionSort(t *testing.T) {
	arr := []int{5, 3, 6, 2, 10}
	wantArr := []int{2, 3, 5, 6, 10}

	gotArr := SelectionSort(arr)

	if SlicesEqual(gotArr, wantArr) {
		t.Errorf("Selection Sort = %v, wanted = %v", gotArr, wantArr)
	}
}
