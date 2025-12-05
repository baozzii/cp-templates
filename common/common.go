package common

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type RealNumber interface {
	Integer |
		~float32 | ~float64
}

func abs[T RealNumber](x T) T {
	if x < T(0) {
		return -x
	}
	return x
}

func gcd[T Integer](x, y T) T {
	if x < 0 || y < 0 {
		return gcd(abs(x), abs(y))
	}
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func lcm[T Integer](x, y T) T {
	return x / gcd(x, y) * y
}

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
