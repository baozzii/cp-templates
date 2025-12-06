package datastructures

type DynamicLazySegtreeInfo[T, K any] interface {
	op(T, T) T
	e() T
	mp(K, T, int) T
	cp(K, K) K
	id() K
}

type DynamicLazySegtree[T, K any, M DynamicLazySegtreeInfo[T, K]] struct {
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

func (seg *DynamicLazySegtree[T, K, M]) pushup(i int) {
	seg.d[i] = seg.op(seg.d[seg.ls[i]], seg.d[seg.rs[i]])
}

func (seg *DynamicLazySegtree[T, K, M]) do_apply(i int, tag K, length int) {
	if i != 0 {
		seg.d[i] = seg.mp(tag, seg.d[i], length)
		seg.lz[i] = seg.cp(tag, seg.lz[i])
	}
}

func (seg *DynamicLazySegtree[T, K, M]) pushdown(i, l, r int) {
	if seg.ls[i] == 0 {
		seg.ls[i] = seg.new_node()
	}
	if seg.rs[i] == 0 {
		seg.rs[i] = seg.new_node()
	}
	m := (l + r) / 2
	seg.do_apply(seg.ls[i], seg.lz[i], m-l)
	seg.do_apply(seg.rs[i], seg.lz[i], r-m)
	seg.lz[i] = seg.id()
}

func (seg *DynamicLazySegtree[T, K, M]) new_node() int {
	seg.d = append(seg.d, seg.e())
	seg.lz = append(seg.lz, seg.id())
	seg.ls = append(seg.ls, 0)
	seg.rs = append(seg.rs, 0)
	return len(seg.d) - 1
}

func NewDynamicLazySegtree[T, K any, M DynamicLazySegtreeInfo[T, K]](lb, rb int, m M) *DynamicLazySegtree[T, K, M] {
	d := make([]T, 2)
	d[0] = m.e()
	d[1] = m.e()
	lz := make([]K, 2)
	lz[0] = m.id()
	lz[1] = m.id()
	return &DynamicLazySegtree[T, K, M]{
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

func (seg *DynamicLazySegtree[T, K, M]) Set(i int, x T) {
	var dfs func(int, int, int)
	dfs = func(o, l, r int) {
		if l+1 == r {
			seg.d[o] = x
			return
		}
		seg.pushdown(o, l, r)
		m := (l + r) / 2
		if i < m {
			if seg.ls[o] == 0 {
				seg.ls[o] = seg.new_node()
			}
			dfs(seg.ls[o], l, m)
		} else {
			if seg.rs[o] == 0 {
				seg.rs[o] = seg.new_node()
			}
			dfs(seg.rs[o], m, r)
		}
		seg.pushup(o)
	}
	dfs(1, seg.lb, seg.rb)
}

func (seg *DynamicLazySegtree[T, K, M]) Apply(ql, qr int, tag K) {
	var dfs func(int, int, int)
	dfs = func(o, l, r int) {
		if o == 0 || r <= ql || qr <= l {
			return
		}
		if ql <= l && r <= qr {
			seg.do_apply(o, tag, r-l)
			return
		}
		seg.pushdown(o, l, r)
		m := (l + r) / 2
		dfs(seg.ls[o], l, m)
		dfs(seg.rs[o], m, r)
		seg.pushup(o)
	}
	dfs(1, seg.lb, seg.rb)
}

func (seg *DynamicLazySegtree[T, K, M]) Prod(ql, qr int) T {
	var dfs func(int, int, int) T
	dfs = func(o, l, r int) T {
		if o == 0 || r <= ql || qr <= l {
			return seg.e()
		}
		if ql <= l && r <= qr {
			return seg.d[o]
		}
		seg.pushdown(o, l, r)
		m := (l + r) / 2
		return seg.op(dfs(seg.ls[o], l, m), dfs(seg.rs[o], m, r))
	}
	return dfs(1, seg.lb, seg.rb)
}

func (seg *DynamicLazySegtree[T, K, M]) Get(i int) T {
	var dfs func(int, int, int) T
	dfs = func(o, l, r int) T {
		if o == 0 {
			return seg.e()
		}
		if l+1 == r {
			return seg.d[o]
		}
		seg.pushdown(o, l, r)
		m := (l + r) / 2
		if i < m {
			return dfs(seg.ls[o], l, m)
		} else {
			return dfs(seg.rs[o], m, r)
		}
	}
	return dfs(1, seg.lb, seg.rb)
}
