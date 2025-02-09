package utils

func GetPointer[T any](elem T) *T {
	return &elem
}
