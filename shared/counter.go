package shared

func MakeCounter[T comparable](arr []T) map[T]int {
	counter := make(map[T]int)
	for _, num := range arr {
		if _, ok := counter[num]; !ok {
			counter[num] = 1
		} else {
			counter[num]++
		}
	}

	return counter
}
