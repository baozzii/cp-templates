package templates

/*
https://judge.yosupo.jp/submission/334845
*/

type dynamic_lazysegtree_info[T, K any] interface {
	op(T, T) T
	e() T
	mp(K, T, int) T
	cp(K, K) K
	id() K
}

type dynamic_lazysegtree[T, K any, M dynamic_lazysegtree_info[T, K]] struct {
	d      []T
	lz     []K
	ls, rs []int
	op     func(T, T) T
	e      func() T
	mp     func(K, T, int) T
	cp     func(K, K) K
	id     func() K
	lb, rb int
}

func (seg *dynamic_lazysegtree[T, K, M]) __pushup(i int) {
	seg.d[i] = seg.op(seg.d[seg.ls[i]], seg.d[seg.rs[i]])
}

func (seg *dynamic_lazysegtree[T, K, M]) __do_apply(i int, tag K, length int) {
	if i != 0 {
		seg.d[i] = seg.mp(tag, seg.d[i], length)
		seg.lz[i] = seg.cp(tag, seg.lz[i])
	}
}

func (seg *dynamic_lazysegtree[T, K, M]) __pushdown(i, l, r int) {
	if seg.ls[i] == 0 {
		seg.ls[i] = seg.__new_node()
	}
	if seg.rs[i] == 0 {
		seg.rs[i] = seg.__new_node()
	}
	m := (l + r) / 2
	seg.__do_apply(seg.ls[i], seg.lz[i], m-l)
	seg.__do_apply(seg.rs[i], seg.lz[i], r-m)
	seg.lz[i] = seg.id()
}

func (seg *dynamic_lazysegtree[T, K, M]) __new_node() int {
	seg.d = append(seg.d, seg.e())
	seg.lz = append(seg.lz, seg.id())
	seg.ls = append(seg.ls, 0)
	seg.rs = append(seg.rs, 0)
	return len(seg.d) - 1
}

func new_dynamic_lazysegtree[T, K any, M dynamic_lazysegtree_info[T, K]](lb, rb int, m M) *dynamic_lazysegtree[T, K, M] {
	d := make([]T, 2)
	d[0] = m.e()
	d[1] = m.e()
	lz := make([]K, 2)
	lz[0] = m.id()
	lz[1] = m.id()
	return &dynamic_lazysegtree[T, K, M]{
		d:  d,
		lz: lz,
		ls: make([]int, 2),
		rs: make([]int, 2),
		op: m.op,
		e:  m.e,
		mp: m.mp,
		cp: m.cp,
		id: m.id,
		lb: lb,
		rb: rb,
	}
}

func (seg *dynamic_lazysegtree[T, K, M]) set(i int, x T) {
	var dfs func(int, int, int)
	dfs = func(o, l, r int) {
		if l+1 == r {
			seg.d[o] = x
			return
		}
		seg.__pushdown(o, l, r)
		m := (l + r) / 2
		if i < m {
			if seg.ls[o] == 0 {
				seg.ls[o] = seg.__new_node()
			}
			dfs(seg.ls[o], l, m)
		} else {
			if seg.rs[o] == 0 {
				seg.rs[o] = seg.__new_node()
			}
			dfs(seg.rs[o], m, r)
		}
		seg.__pushup(o)
	}
	dfs(1, seg.lb, seg.rb)
}

func (seg *dynamic_lazysegtree[T, K, M]) apply(ql, qr int, tag K) {
	var dfs func(int, int, int)
	dfs = func(o, l, r int) {
		if o == 0 || r <= ql || qr <= l {
			return
		}
		if ql <= l && r <= qr {
			seg.__do_apply(o, tag, r-l)
			return
		}
		seg.__pushdown(o, l, r)
		m := (l + r) / 2
		dfs(seg.ls[o], l, m)
		dfs(seg.rs[o], m, r)
		seg.__pushup(o)
	}
	dfs(1, seg.lb, seg.rb)
}

func (seg *dynamic_lazysegtree[T, K, M]) prod(ql, qr int) T {
	var dfs func(int, int, int) T
	dfs = func(o, l, r int) T {
		if o == 0 || r <= ql || qr <= l {
			return seg.e()
		}
		if ql <= l && r <= qr {
			return seg.d[o]
		}
		seg.__pushdown(o, l, r)
		m := (l + r) / 2
		return seg.op(dfs(seg.ls[o], l, m), dfs(seg.rs[o], m, r))
	}
	return dfs(1, seg.lb, seg.rb)
}

func (seg *dynamic_lazysegtree[T, K, M]) get(i int) T {
	var dfs func(int, int, int) T
	dfs = func(o, l, r int) T {
		if o == 0 {
			return seg.e()
		}
		if l+1 == r {
			return seg.d[o]
		}
		seg.__pushdown(o, l, r)
		m := (l + r) / 2
		if i < m {
			return dfs(seg.ls[o], l, m)
		} else {
			return dfs(seg.rs[o], m, r)
		}
	}
	return dfs(1, seg.lb, seg.rb)
}
