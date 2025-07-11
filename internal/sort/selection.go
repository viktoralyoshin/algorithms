package sort

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIdx := 0

	for idx, value := range arr {
		if value < smallest {
			smallest = value
			smallestIdx = idx
		}
	}

	return smallestIdx
}

func SelectionSort(arr []int) []int {
	newArr := make([]int, 0, len(arr))

	cpArr := make([]int, len(arr))
	copy(cpArr, arr)

	for range cpArr {
		smallest := findSmallest(cpArr)

		newArr = append(newArr, cpArr[smallest])
		cpArr = append(cpArr[:smallest], cpArr[smallest+1:]...)
	}

	return newArr
}
