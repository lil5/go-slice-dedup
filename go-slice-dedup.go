package goslicededup

func DeDuplicate[T string | int | int32 | int64 | float32 | float64 | uint](list []T) []T {
	keys := map[T]bool{}
	newList := []T{}
	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range list {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			newList = append(newList, entry)
		}
	}
	return newList
}
