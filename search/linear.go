package search

func Linear(arr []int, target int) (int, error) {
	for i, item := range arr {
		if item == target {
			return i, nil
		}
	}
	return -1, ErrNotFound
}
