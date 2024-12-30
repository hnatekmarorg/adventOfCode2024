package utils

func CountOccurrences[T comparable](val T, slice []T) int {
	count := 0
	for _, v := range slice {
		if v == val {
			count++
		}
	}
	return count
}
