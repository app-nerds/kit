package slices

/*
IsInSlice checks if an item is in a slice. This function will work
on anything that implements the comparable interface.
*/
func IsInSlice[T comparable](item T, slice []T) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}

	return false
}
