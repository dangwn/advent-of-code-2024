package dayone

import "strconv"

func absDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func convArr(arr []string) ([]int, error) {
	var out []int
	for _, val := range arr {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return out, err
		}
		out = append(out, intVal)
	}
	return out, nil
}

func makeCounter(arr []int) map[int]int {
	counter := make(map[int]int)
	for _, num := range arr {
		if _, ok := counter[num]; !ok {
			counter[num] = 1
		} else {
			counter[num]++
		}
	}

	return counter
}
