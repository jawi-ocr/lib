package number

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float | constraints.Unsigned
}

type Number[T Numeric] struct {
	Value T
}

func (n Number[T]) String() string {
	return fmt.Sprintf("%v", n.Value)
}

func New[T Numeric](number T) Number[T] {
	return Number[T]{Value: number}
}

type Numbers[T Numeric] []Number[T]

func (n *Numbers[T]) Add(number T) {
	*n = append(*n, New[T](number))
}

func NewSlice[T Numeric]() Numbers[T] {
	return make([]Number[T], 0)
}