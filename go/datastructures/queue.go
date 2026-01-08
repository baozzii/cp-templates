package datastructures

type Queue[T any] struct {
	t []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Push(x T) {
	q.t = append(q.t, x)
}

func (q *Queue[T]) Front() T {
	return q.t[0]
}

func (q *Queue[T]) Pop() {
	q.t = q.t[1:]
}

func (q *Queue[T]) Empty() bool {
	return len(q.t) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.t)
}
