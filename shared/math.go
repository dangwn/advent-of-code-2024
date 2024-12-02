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
