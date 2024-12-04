package shared

func Any[T any](data []T, cond func(T) bool) bool {
	for _, d := range data {
		if cond(d) {
			return true
		}
	}
	return false
}

func All[T any](data []T, cond func(T) bool) bool {
	for _, d := range data {
		if !cond(d) {
			return false
		}
	}
	return true
}

func Map[T, V any](data []T, mapper func(T) V) []V {
	arr := make([]V, 0, len(data))
	for _, d := range data {
		arr = append(arr, mapper(d))
	}
	return arr
}

func FirstInd[T any](data []T, cond func(T) bool) int {
	for ind, d := range data {
		if cond(d) {
			return ind
		}
	}
	return -1
}

func AllTrue(cond []bool) bool {
	for _, c := range cond {
		if !c {
			return false
		}
	}
	return true
}
