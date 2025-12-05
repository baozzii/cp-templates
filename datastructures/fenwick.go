package datastructures

type FenwickInfo[T any] interface {
	add(T, T) T
	sub(T, T) T
	e() T
}

type Fenwick[T any, M FenwickInfo[T]] struct {
	t   []T
	n   int
	add func(T, T) T
	sub func(T, T) T
	e   func() T
}

func NewFenwick[T any, M FenwickInfo[T]](n int, m M) *Fenwick[T, M] {
	t := make([]T, n+1)
	for i := range t {
		t[i] = m.e()
	}
	return &Fenwick[T, M]{t, n, m.add, m.sub, m.e}
}

func (fen *Fenwick[T, M]) Add(i int, x T) {
	for i++; i <= fen.n; i += i & -i {
		fen.t[i] = fen.add(fen.t[i], x)
	}
}

func (fen *Fenwick[T, M]) Pre(i int) T {
	add := fen.add
	r := fen.e()
	for ; i > 0; i &= i - 1 {
		r = add(r, fen.t[i])
	}
	return r
}

func (fen *Fenwick[T, M]) Sum(l, r int) T {
	return fen.sub(fen.Pre(r), fen.Pre(l))
}
