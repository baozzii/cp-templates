package common

import (
	"slices"
)

type Vec[T any] []T

func (v *Vec[T]) Size() int {
	return len(*v)
}

func (v *Vec[T]) Empty() bool {
	return len(*v) == 0
}

func (v *Vec[T]) Clear() {
	*v = (*v)[:0]
}

func (v *Vec[T]) PushBack(x T) {
	*v = append(*v, x)
}

func (v *Vec[T]) PopBack() {
	*v = (*v)[:len(*v)-1]
}

func (v *Vec[T]) Resize(x int) {
	if v.Size() > x {
		*v = (*v)[:x]
	} else {
		*v = append(*v, make([]T, x-v.Size())...)
	}
}

func (v *Vec[T]) Back() *T {
	return &((*v)[v.Size()-1])
}

func (v *Vec[T]) Front() *T {
	return &((*v)[0])
}

func (v *Vec[T]) Erase(x, y int) {
	*v = slices.Delete(*v, x, y)
}

func (v *Vec[T]) Insert(x int, w ...T) {
	*v = slices.Insert(*v, x, w...)
}

func (v *Vec[T]) Reverse() {
	n := v.Size()
	for i := 0; i < n/2; i++ {
		(*v)[i], (*v)[n-i-1] = (*v)[n-i-1], (*v)[i]
	}
}

func (v *Vec[T]) Fill(w T) {
	for i := range *v {
		(*v)[i] = w
	}
}

func ToVec[T any](v []T) *Vec[T] {
	t := Vec[T](v)
	return &t
}

func Vec1[T any](n int) *Vec[T] {
	v := make(Vec[T], n)
	return &v
}

func Vec2[T any](n, m int) *Vec[Vec[T]] {
	v := make(Vec[Vec[T]], n)
	for i := range v {
		v[i] = make(Vec[T], m)
	}
	return &v
}

func (v *Vec[T]) Copy() *Vec[T] {
	nv := slices.Clone(*v)
	return &nv
}
