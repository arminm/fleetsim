package common

func RemoveItemByReference[T any](slice []T, target T) []T {
	for i, item := range slice {
		if &item == &target {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
