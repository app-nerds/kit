package slices

/*
Filter applies the function 'test' to all values in a slice and returns
a new slice for all values where 'test' returned true.
*/
func Filter[T any](a []T, test func(T) bool) []T {
	result := make([]T, 0, len(a))

	for _, value := range a {
		if test(value) {
			result = append(result, value)
		}
	}

	return result
}
