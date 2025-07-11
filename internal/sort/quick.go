package sort

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		pivot := arr[0]
		var less, greater []int

		for _, value := range arr[1:] {
			if value <= pivot {
				less = append(less, value)
			} else {
				greater = append(greater, value)
			}
		}
		
		less = QuickSort(less)
		greater = QuickSort(greater)
		
		result := append(less, pivot)
		result = append(result, greater...)
		
		return result
	}
}
