package utils

func HasDuplicates[T comparable](slice []T) bool {
	seen := map[T]bool{}
	for _, item := range slice {
		if _, ok := seen[item]; ok {
			return true
		}
		seen[item] = true
	}
	return false
}
