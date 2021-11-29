package sorts

func selection_sort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
