package datastructures

import . "codeforces-go/common"

type SegtreeInfo[T any] interface {
	op(T, T) T
	e() T
}

type Segtree[T any, M SegtreeInfo[T]] struct {
	n, size, log int
	d            []T
	op           func(T, T) T
	e            func() T
}

func (seg *Segtree[T, M]) pushup(i int) {
	seg.d[i] = seg.op(seg.d[i<<1], seg.d[i<<1|1])
}

func NewSegtreeWith[T any, M SegtreeInfo[T]](a []T, m M) *Segtree[T, M] {
	size := 1
	n := len(a)
	for size < n {
		size <<= 1
	}
	log := Ctz(size)
	d := make([]T, size<<1)
	for i := range d {
		d[i] = m.e()
	}
	for i := range a {
		d[i+size] = a[i]
	}
	for i := size - 1; i > 0; i-- {
		d[i] = m.op(d[i<<1], d[i<<1|1])
	}
	return &Segtree[T, M]{n, size, log, d, m.op, m.e}
}

func NewSegtree[T any, M SegtreeInfo[T]](n int, m M) *Segtree[T, M] {
	a := make([]T, n)
	for i := range a {
		a[i] = m.e()
	}
	return NewSegtreeWith(a, m)
}

func (seg *Segtree[T, M]) Set(p int, x T) {
	p += seg.size
	seg.d[p] = x
	for i := 1; i <= seg.log; i++ {
		seg.pushup(p >> i)
	}
}

func (seg *Segtree[T, M]) Get(p int) T {
	return seg.d[p+seg.size]
}

func (seg *Segtree[T, M]) Prod(l, r int) T {
	e := seg.e
	op := seg.op
	sml := e()
	smr := e()
	for l, r = l+seg.size, r+seg.size; l < r; l, r = l>>1, r>>1 {
		if l%2 == 1 {
			sml = op(sml, seg.d[l])
			l++
		}
		if r%2 == 1 {
			r--
			smr = op(seg.d[r], smr)
		}
	}
	return op(sml, smr)
}

func (seg *Segtree[T, M]) AllProd() T {
	return seg.d[1]
}

func (seg *Segtree[T, M]) MaxRight(l int, f func(T) bool) int {
	if l == seg.n {
		return seg.n
	}
	e := seg.e
	op := seg.op
	l += seg.size
	sm := e()
	for i := 0; (l&-l) != l || i == 0; i++ {
		for l%2 == 0 {
			l >>= 1
		}
		if !f(op(sm, seg.d[l])) {
			for l < seg.size {
				l <<= 1
				if f(op(sm, seg.d[l])) {
					sm = op(sm, seg.d[l])
					l++
				}
			}
			return l - seg.size
		}
		sm = op(sm, seg.d[l])
		l++
	}
	return seg.n
}

func (seg *Segtree[T, M]) MinLeft(r int, f func(T) bool) int {
	if r == 0 {
		return 0
	}
	e := seg.e
	op := seg.op
	r += seg.size
	sm := e()
	for i := 0; (r&-r) != r || i == 0; i++ {
		r--
		for r > 1 && r%2 == 1 {
			r >>= 1
		}
		if !f(op(seg.d[r], sm)) {
			for r < seg.size {
				r = r<<1 | 1
				if f(op(seg.d[r], sm)) {
					sm = op(seg.d[r], sm)
					r--
				}
			}
			return r + 1 - seg.size
		}
		sm = op(seg.d[r], sm)
	}
	return 0
}

type SegMax struct{}

func (SegMax) op(x, y int) int { return max(x, y) }
func (SegMax) e() int          { return int(-1e18) }

type SegMin struct{}

func (SegMin) op(x, y int) int { return min(x, y) }
func (SegMin) e() int          { return int(1e18) }

type SegSum struct{}

func (SegSum) op(x, y int) int { return x + y }
func (SegSum) e() int          { return 0 }
