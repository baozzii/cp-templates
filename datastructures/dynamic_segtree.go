package datastructures

type DynamicSegtreeInfo[T any] interface {
	op(T, T) T
	e() T
}

type DynamicSegtree[T any, M DynamicSegtreeInfo[T]] struct {
	d      []T
	ls, rs []int
	op     func(T, T) T
	e      func() T
	lb, rb int
}

func (seg *DynamicSegtree[T, M]) pushup(i int) {
	seg.d[i] = seg.op(seg.d[seg.ls[i]], seg.d[seg.rs[i]])
}

func (seg *DynamicSegtree[T, M]) new_node() int {
	seg.d = append(seg.d, seg.e())
	seg.ls = append(seg.ls, 0)
	seg.rs = append(seg.rs, 0)
	return len(seg.d) - 1
}

func NewDynamicSegtree[T any, M DynamicSegtreeInfo[T]](lb, rb int, m M) *DynamicSegtree[T, M] {
	d := make([]T, 2)
	d[0] = m.e()
	d[1] = m.e()
	return &DynamicSegtree[T, M]{d, make([]int, 2), make([]int, 2), m.op, m.e, lb, rb}
}

func (seg *DynamicSegtree[T, M]) Set(i int, x T) {
	var dfs func(int, int, int)
	dfs = func(o, l, r int) {
		if l+1 == r {
			seg.d[o] = x
			return
		}
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

func (seg *DynamicSegtree[T, M]) Prod(ql, qr int) T {
	var dfs func(int, int, int) T
	dfs = func(o, l, r int) T {
		if o == 0 || r <= ql || qr <= l {
			return seg.e()
		}
		if ql <= l && qr >= r {
			return seg.d[o]
		}
		m := (l + r) / 2
		return seg.op(dfs(seg.ls[o], l, m), dfs(seg.rs[o], m, r))
	}
	return dfs(1, seg.lb, seg.rb)
}

func (seg *DynamicSegtree[T, M]) Get(i int) T {
	var dfs func(int, int, int) T
	dfs = func(o, l, r int) T {
		if o == 0 {
			return seg.e()
		}
		if l+1 == r {
			return seg.d[o]
		}
		m := (l + r) / 2
		if i < m {
			return dfs(seg.ls[o], l, m)
		} else {
			return dfs(seg.rs[o], m, r)
		}
	}
	return dfs(1, seg.lb, seg.rb)
}
