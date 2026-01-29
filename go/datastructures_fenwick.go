package templates

/*
https://judge.yosupo.jp/submission/334041
*/

type fenwick_info[T any] interface {
	add(T, T) T
	sub(T, T) T
	e() T
}

type fenwick[T any, M fenwick_info[T]] struct {
	t     []T
	n     int
	__add func(T, T) T
	__sub func(T, T) T
	e     func() T
}

func new_fenwick[T any, M fenwick_info[T]](n int, m M) *fenwick[T, M] {
	t := make([]T, n+1)
	for i := range t {
		t[i] = m.e()
	}
	return &fenwick[T, M]{t, n, m.add, m.sub, m.e}
}

func (fen *fenwick[T, M]) add(i int, x T) {
	for i++; i <= fen.n; i += i & -i {
		fen.t[i] = fen.__add(fen.t[i], x)
	}
}

func (fen *fenwick[T, M]) pre(i int) T {
	r := fen.e()
	for ; i > 0; i &= i - 1 {
		r = fen.__add(r, fen.t[i])
	}
	return r
}

func (fen *fenwick[T, M]) Sum(l, r int) T {
	return fen.__sub(fen.pre(r), fen.pre(l))
}

func (fen *fenwick[T, M]) kth(k T, cmp func(x, y T) int) int {
	u := 0
	for d := highbit(fen.n) + 1; d >= 0; d-- {
		if u+(1<<d) <= fen.n && cmp(k, fen.t[u+(1<<d)]) >= 0 {
			u += 1 << d
			k = fen.__sub(k, fen.t[u+(1<<d)])
		}
	}
	return u
}

type fen_sum struct{}

func (fen_sum) add(x, y int) int { return x + y }
func (fen_sum) sub(x, y int) int { return x - y }
func (fen_sum) e() int           { return 0 }

type fen_max struct{}

func (fen_max) add(x, y int) int { return max(x, y) }
func (fen_max) e() int           { return int(-1e18) }

type fen_xor struct{}

func (fen_xor) add(x, y int) int { return x ^ y }
func (fen_xor) sub(x, y int) int { return x ^ y }
func (fen_xor) e() int           { return 0 }
