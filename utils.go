package playstation

func IsContain[T comparable](arr []T, value T) bool {
	for _, elem := range arr {
		if elem == value {
			return true
		}
	}
	return false
}

func must[T any](ret T, err error) T {
	if err != nil {
		panic(err)
	}
	return ret
}
