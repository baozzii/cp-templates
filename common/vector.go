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

func ToVec1[T any, E ~[]T](v E) Vec[T] {
	return Vec[T](v)
}

func ToSlice1[T any](v Vec[T]) []T {
	return []T(v)
}

func ToVec2[T any, E ~[][]T](v E) Vec[Vec[T]] {
	res := make(Vec[Vec[T]], len(v))
	for i := range v {
		res[i] = ToVec1(v[i])
	}
	return res
}

func ToSlice2[T any](v Vec[Vec[T]]) [][]T {
	res := make([][]T, len(v))
	for i := range v {
		res[i] = ToSlice1(v[i])
	}
	return res
}

func ToVec3[T any, E ~[][][]T](v E) Vec[Vec[Vec[T]]] {
	res := make(Vec[Vec[Vec[T]]], len(v))
	for i := range v {
		res[i] = ToVec2(v[i])
	}
	return res
}

func ToSlice3[T any](v Vec[Vec[Vec[T]]]) [][][]T {
	res := make([][][]T, len(v))
	for i := range v {
		res[i] = ToSlice2(v[i])
	}
	return res
}

func Vec1[T any](n ...int) Vec[T] {
	if len(n) == 0 {
		return make(Vec[T], 0)
	}
	return make(Vec[T], n[0])
}

func Vec2[T any](n ...int) Vec[Vec[T]] {
	if len(n) == 0 {
		return make(Vec[Vec[T]], 0)
	}
	v := make(Vec[Vec[T]], n[0])
	for i := range v {
		v[i] = Vec1[T](n[1:]...)
	}
	return v
}

func Vec3[T any](n ...int) Vec[Vec[Vec[T]]] {
	if len(n) == 0 {
		return make(Vec[Vec[Vec[T]]], 0)
	}
	v := make(Vec[Vec[Vec[T]]], n[0])
	for i := range v {
		v[i] = Vec2[T](n[1:]...)
	}
	return v
}

func (v *Vec[T]) Copy() Vec[T] {
	return slices.Clone(*v)
}
