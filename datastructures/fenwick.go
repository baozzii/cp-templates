package datastructures

type fenwick_info[T any] interface {
	add(T, T) T
	sub(T, T) T
	e() T
}

type fenwick[T any, M fenwick_info[T]] struct {
	t   []T
	n   int
	add func(T, T) T
	sub func(T, T) T
	e   func() T
}

func new_fenwick[T any, M fenwick_info[T]](n int, m M) *fenwick[T, M] {
	t := make([]T, n+1)
	for i := range t {
		t[i] = m.e()
	}
	return &fenwick[T, M]{t, n, m.add, m.sub, m.e}
}

func (fen *fenwick[T, M]) inc(i int, x T) {
	for i++; i <= fen.n; i += i & -i {
		fen.t[i] = fen.add(fen.t[i], x)
	}
}

func (fen *fenwick[T, M]) pre(i int) T {
	add := fen.add
	r := fen.e()
	for ; i > 0; i &= i - 1 {
		r = add(r, fen.t[i])
	}
	return r
}

func (fen *fenwick[T, M]) sum(l, r int) T {
	return fen.sub(fen.pre(r), fen.pre(l))
}
