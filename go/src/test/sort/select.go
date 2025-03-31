package sort

func selectSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		min := i
		for j := 1 + i; j < n; j++ {
			if arr[min] > arr[j] {
				min = j
			}

		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
