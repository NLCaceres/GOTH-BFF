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

// Keeps unique elements based on a function using the slice element to derive a selector value
// for comparison of later elements to prevent duplicates in the returned slice
func DistinctBy[T any, U comparable](list []T, selector func(T) U) []T {
	valueMap := make(map[U]bool)
	newList := make([]T, 0, len(list))
	for _, value := range list {
		if !valueMap[selector(value)] {
			valueMap[selector(value)] = true
			newList = append(newList, value)
		}
	}

	return newList
}
