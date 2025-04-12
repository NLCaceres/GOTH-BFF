package list

// Keep elements based on validation function and return these elements in new slice
func Filter[T any](list []T, valid func(T) bool) []T {
	newList := make([]T, 0, len(list))
	for _, value := range list {
		if valid(value) { // Must append() the kept/valid elements to increase new slice length
			newList = append(newList, value)
		}
	}
	return newList
}
