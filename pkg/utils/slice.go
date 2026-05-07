// pkg/utils/slice.go

package utils

func Filter[T any](data []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range data {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map[T any, U any](data []T, transform func(T) U) []U {
	var result []U
	for _, item := range data {
		result = append(result, transform(item))
	}
	return result
}
