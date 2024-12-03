package shared

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Abs[T Number](val T) T {
	if val < 0 {
		return -val
	}
	return val
}

func Sum[T Number](vals []T) T {
	var total T = 0
	for _, v := range vals {
		total += v
	}
	return total
}
