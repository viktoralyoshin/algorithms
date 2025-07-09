package search

import (
	"fmt"
)

func BinarySearch(slc []int, item int) (int, error){
	l := 0
	r := len(slc)-1

	for l<=r {
		m := (l+r)/2

		if slc[m] == item {
			return m, nil
		}

		if slc[m] > item {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1, fmt.Errorf("element %d not found in slice", item)
}

