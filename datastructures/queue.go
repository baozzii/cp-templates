package datastructures

type queue[T any] struct {
	t []T
}

func new_queue[T any]() *queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) push(x T) {
	q.t = append(q.t, x)
}

func (q *queue[T]) front() T {
	return q.t[0]
}

func (q *queue[T]) pop() {
	q.t = q.t[1:]
}
