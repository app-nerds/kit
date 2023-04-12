package slices

/*
BreakIntoGroups breaks a slice into groups of a given size.
*/
func BreakIntoGroups[T any](collection []T, groupSize int) [][]T {
	var (
		result [][]T
		end    int
	)

	for i := 0; i < len(collection); i += groupSize {
		end += i + groupSize

		if end > len(collection) {
			end = len(collection)
		}

		result = append(result, collection[i:end])
	}

	/* We have to convert the result into a new slice of type T */
	newSlice := make([][]T, len(result))

	for i, r := range result {
		temp := make([]T, len(r))

		copy(temp, r)
		newSlice[i] = temp
	}

	return newSlice
}
