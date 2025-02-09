package utils

func HasDuplicatesMapped[S ~[]E, E any, M comparable](slice S, mapFunc func(index int, elem E) M) bool {
	seen := map[M]bool{}
	for index, item := range slice {
		mappedItem := mapFunc(index, item)
		if _, ok := seen[mappedItem]; ok {
			return true
		}
		seen[mappedItem] = true
	}
	return false
}

func HasDuplicates[T comparable](slice []T) bool {
	return HasDuplicatesMapped(slice, func(index int, item T) T {
		return item
	})
}
