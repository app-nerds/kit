package slices

/*
ConvertSlice converts all the values in slice 'a' from type
T to type K using a converter function.

	converted := slices.ConvertSlice[graphql.String, string](someStruct.GraphQLStrings, func(input graphql.String) string {
	  return string(input)
	})
*/
func ConvertSlice[T any, K any](a []T, convertFunc func(T) K) []K {
	result := make([]K, len(a))

	for index, value := range a {
		result[index] = convertFunc(value)
	}

	return result
}
