package templates

type lazysegtree_info[T, K any] interface {
	op(T, T) T
	e() T
	mp(K, T) T
	cp(K, K) K
	id() K
}

type lazysegtree[T, K any, M lazysegtree_info[T, K]] struct {
	n, size, log int
	d            []T
	lz           []K
	op           func(T, T) T
	e            func() T
	mp           func(K, T) T
	cp           func(K, K) K
	id           func() K
}

func (seg *lazysegtree[T, K, M]) __pushup(k int) {
	seg.d[k] = seg.op(seg.d[k<<1], seg.d[k<<1|1])
}

func (seg *lazysegtree[T, K, M]) __all_apply(k int, f K) {
	seg.d[k] = seg.mp(f, seg.d[k])
	if k < seg.size {
		seg.lz[k] = seg.cp(f, seg.lz[k])
	}
}

func (seg *lazysegtree[T, K, M]) __push(k int) {
	f := seg.lz[k]
	if k < seg.size {
		seg.__all_apply(k<<1, f)
		seg.__all_apply(k<<1|1, f)
	}
	seg.lz[k] = seg.id()
}

func new_lazysegtree_with[T, K any, M lazysegtree_info[T, K], E ~[]T](a E, m M) *lazysegtree[T, K, M] {
	n := len(a)
	size := 1
	for size < n {
		size <<= 1
	}
	log := ctz(size)
	op := m.op
	e := m.e
	mp := m.mp
	cp := m.cp
	id := m.id
	d := make([]T, size<<1)
	lz := make([]K, size)
	for i := range d {
		d[i] = e()
	}
	for i := range lz {
		lz[i] = id()
	}

	for i := 0; i < n; i++ {
		d[size+i] = a[i]
	}
	for i := size - 1; i > 0; i-- {
		d[i] = op(d[i<<1], d[i<<1|1])
	}

	return &lazysegtree[T, K, M]{
		n:    n,
		size: size,
		log:  log,
		d:    d,
		lz:   lz,
		op:   op,
		e:    e,
		mp:   mp,
		cp:   cp,
		id:   id,
	}
}

func new_lazy_segtree[T, K any, M lazysegtree_info[T, K]](n int, m M) *lazysegtree[T, K, M] {
	a := make([]T, n)
	e := m.e
	for i := range a {
		a[i] = e()
	}
	return new_lazysegtree_with(a, m)
}

func (seg *lazysegtree[T, K, M]) set(p int, x T) {
	p += seg.size
	for i := seg.log; i > 0; i-- {
		seg.__push(p >> i)
	}
	seg.d[p] = x
	for i := 1; i <= seg.log; i++ {
		seg.__pushup(p >> i)
	}
}

func (seg *lazysegtree[T, K, M]) get(p int) T {
	p += seg.size
	for i := seg.log; i > 0; i-- {
		seg.__push(p >> i)
	}
	return seg.d[p]
}

func (seg *lazysegtree[T, K, M]) prod(l, r int) T {
	if l == r {
		return seg.e()
	}
	op := seg.op
	e := seg.e

	l += seg.size
	r += seg.size

	for i := seg.log; i > 0; i-- {
		if ((l >> i) << i) != l {
			seg.__push(l >> i)
		}
		if ((r >> i) << i) != r {
			seg.__push((r - 1) >> i)
		}
	}

	sml := e()
	smr := e()
	for l < r {
		if l&1 == 1 {
			sml = op(sml, seg.d[l])
			l++
		}
		if r&1 == 1 {
			r--
			smr = op(seg.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return op(sml, smr)
}

func (seg *lazysegtree[T, K, M]) all_prod() T {
	return seg.d[1]
}

func (seg *lazysegtree[T, K, M]) apply(l, r int, f K) {
	if l == r {
		return
	}
	l0 := l + seg.size
	r0 := r + seg.size
	for i := seg.log; i > 0; i-- {
		if ((l0 >> i) << i) != l0 {
			seg.__push(l0 >> i)
		}
		if ((r0 >> i) << i) != r0 {
			seg.__push((r0 - 1) >> i)
		}
	}
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			seg.__all_apply(l, f)
			l++
		}
		if r&1 == 1 {
			r--
			seg.__all_apply(r, f)
		}
		l >>= 1
		r >>= 1
	}
	l = l0
	r = r0
	for i := 1; i <= seg.log; i++ {
		if ((l >> i) << i) != l {
			seg.__pushup(l >> i)
		}
		if ((r >> i) << i) != r {
			seg.__pushup((r - 1) >> i)
		}
	}
}

func (seg *lazysegtree[T, K, M]) max_right(l int, f func(T) bool) int {
	if l == seg.n {
		return seg.n
	}
	op := seg.op
	e := seg.e
	l += seg.size
	for i := seg.log; i > 0; i-- {
		seg.__push(l >> i)
	}
	sm := e()
	for {
		for l&1 == 0 {
			l >>= 1
		}
		if !f(op(sm, seg.d[l])) {
			for l < seg.size {
				seg.__push(l)
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
		if (l & -l) == l {
			break
		}
	}
	return seg.n
}

func (seg *lazysegtree[T, K, M]) min_left(r int, f func(T) bool) int {
	if r == 0 {
		return 0
	}
	op := seg.op
	e := seg.e
	r += seg.size
	for i := seg.log; i > 0; i-- {
		seg.__push((r - 1) >> i)
	}
	sm := e()
	for {
		r--
		for r > 1 && r&1 == 1 {
			r >>= 1
		}
		if !f(op(seg.d[r], sm)) {
			for r < seg.size {
				seg.__push(r)
				r = r<<1 | 1
				if f(op(seg.d[r], sm)) {
					sm = op(seg.d[r], sm)
					r--
				}
			}
			return r + 1 - seg.size
		}
		sm = op(seg.d[r], sm)
		if (r & -r) == r {
			break
		}
	}
	return 0
}
