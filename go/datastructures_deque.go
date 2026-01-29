package templates

type deque[T any] struct{ l, r []T }

func (q *deque[T]) empty() bool {
	return len(q.l) == 0 && len(q.r) == 0
}

func (q *deque[T]) size() int {
	return len(q.l) + len(q.r)
}

func (q *deque[T]) push_front(v T) {
	q.l = append(q.l, v)
}

func (q *deque[T]) push_back(v T) {
	q.r = append(q.r, v)
}

func (q *deque[T]) pop_front() (v T) {
	if len(q.l) > 0 {
		q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
	} else {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}

func (q *deque[T]) pop_back() (v T) {
	if len(q.r) > 0 {
		q.r, v = q.r[:len(q.r)-1], q.r[len(q.r)-1]
	} else {
		v, q.l = q.l[0], q.l[1:]
	}
	return
}

func (q *deque[T]) front() *T {
	if len(q.l) > 0 {
		return &q.l[len(q.l)-1]
	}
	return &q.r[0]
}

func (q *deque[T]) back() *T {
	if len(q.r) > 0 {
		return &q.r[len(q.r)-1]
	}
	return &q.l[0]
}

func (q *deque[T]) get(i int) T {
	if i < len(q.l) {
		return q.l[len(q.l)-1-i]
	}
	return q.r[i-len(q.l)]
}
