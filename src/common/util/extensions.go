package util

func Map[T any, U any](items []T, fn func(T) U) []U {
	result := make([]U, len(items))
	for i, v := range items {
		result[i] = fn(v)
	}
	return result
}

func Filter[T any](items []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range items {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func Any[T any](items []T, fn func(T) bool) bool {
	for _, v := range items {
		if fn(v) {
			return true
		}
	}
	return false
}
