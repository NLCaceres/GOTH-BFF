package slice

func SafeMap[T any, U any](list []T, mapper func(T) U) []U {
	newList, _ := ForEach(list, func(t T) (U, error) {
		return mapper(t), nil
	})
	return newList
}

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

// Takes an array and copies it by the number of times input in the "copies" arg,
// turning slices like [1,2] into [1,2,1,2] if duplicated by a factor of 2.
// When "copies" is set to an int below 2, then the slice is simply returned as is.
func Duplicate[T any](s []T, copies int) []T {
	if copies < 2 {
		return s
	}
	list := make([]T, 0, len(s)*copies)
	for range copies {
		list = append(list, s...)
	}
	return list
}
