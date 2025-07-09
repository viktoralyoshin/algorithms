package sort

func BubbleSort(slc []int) []int {
	for i := 0; i < len(slc)-1; i++ {
		for j := i+1; j < len(slc); j++ {
			if slc[i] > slc[j] {
				slc[i], slc[j] = slc[j], slc[i]
			}
		}
	}

	return slc
}
