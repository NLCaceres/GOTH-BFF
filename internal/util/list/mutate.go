package list

// Transform each element of an array/slice via mapping function, then insert them into
// a new slice. If the mapper fails, returns an empty slice with the mapping function error
func ForEach[T any, U any](list []T, mapper func(T) (U, error)) ([]U, error) {
	newList := make([]U, len(list))
	for i, value := range list {
		newValue, err := mapper(value)
		if err != nil {
			return []U{}, err
		}
		newList[i] = newValue
	}
	return newList, nil
}
