package search

import (
	"fmt"
)

func BinarySearch(arr []int, item int) (int, error) {
	l := 0
	r := len(arr) - 1

	for l <= r {
		m := (l + r) / 2

		if arr[m] == item {
			return m, nil
		}

		if arr[m] > item {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1, fmt.Errorf("element %d not found in slice", item)
}
