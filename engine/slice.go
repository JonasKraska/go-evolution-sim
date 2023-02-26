package engine

func PtrSliceRemovceFunc[E comparable](slice []E, callback func(v E) bool) []E {

	index := 0
	for _, element := range slice {
		if callback(element) == true {
			continue
		}

		slice[index] = element
		index++
	}

	// clean up to prevent memory leaks
	//	for i := index; i < len(slice); i++ {
	//		slice[i] = nil
	//	}

	// resize resize slice to remove nil pointers
	return slice[:index]
}

func PtrSliceRemovce[E comparable](slice []E, value E) []E {
	return PtrSliceRemovceFunc(slice, func(v E) bool {
		return v == value
	})
}
