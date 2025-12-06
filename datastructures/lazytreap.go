package datastructures

/*
https://judge.yosupo.jp/submission/334576
*/

import "math/rand"

type LazyTreapInfo[T any, K any] interface {
	op(T, T) T
	e() T
	mp(K, T) T
	cp(K, K) K
	id() K
	rev(T, K) (T, K)
}

type LazyTreapNode[T any, K any, M LazyTreapInfo[T, K]] struct {
	ls, rs, sz int
	pri        uint
	val, sum   T
	tag        K
	rev        bool
}

type LazyTreap[T any, K any, M LazyTreapInfo[T, K]] struct {
	root int
	tr   []LazyTreapNode[T, K, M]
	op   func(T, T) T
	e    func() T
	mp   func(K, T) T
	cp   func(K, K) K
	id   func() K
	rev  func(T, K) (T, K)
}

func NewLazyTreap[T any, K any, M LazyTreapInfo[T, K]](m M) *LazyTreap[T, K, M] {
	tr := make([]LazyTreapNode[T, K, M], 1)
	tr[0] = LazyTreapNode[T, K, M]{0, 0, 0, 0, m.e(), m.e(), m.id(), false}
	return &LazyTreap[T, K, M]{0, tr, m.op, m.e, m.mp, m.cp, m.id, m.rev}
}

func (tr *LazyTreap[T, K, M]) new_node(x T) int {
	tr.tr = append(tr.tr, LazyTreapNode[T, K, M]{0, 0, 1, uint(rand.Uint64()), x, x, tr.id(), false})
	return len(tr.tr) - 1
}

func (tr *LazyTreap[T, K, M]) ls(o int) int {
	return tr.tr[o].ls
}

func (tr *LazyTreap[T, K, M]) rs(o int) int {
	return tr.tr[o].rs
}

func (tr *LazyTreap[T, K, M]) sz(o int) int {
	return tr.tr[o].sz
}

func (tr *LazyTreap[T, K, M]) val(o int) T {
	return tr.tr[o].val
}

func (tr *LazyTreap[T, K, M]) sum(o int) T {
	return tr.tr[o].sum
}

func (tr *LazyTreap[T, K, M]) pri(o int) uint {
	return tr.tr[o].pri
}

func (tr *LazyTreap[T, K, M]) pushup(o int) {
	tr.tr[o].sz = 1 + tr.sz(tr.ls(o)) + tr.sz(tr.rs(o))
	tr.tr[o].sum = tr.op(tr.sum(tr.ls(o)), tr.op(tr.val(o), tr.sum(tr.rs(o))))
}

func (tr *LazyTreap[T, K, M]) apply_tag(o int, tag K) {
	if o == 0 {
		return
	}
	n := &tr.tr[o]
	n.sum = tr.mp(tag, n.sum)
	n.val = tr.mp(tag, n.val)
	n.tag = tr.cp(tag, n.tag)
}

func (tr *LazyTreap[T, K, M]) apply_rev(o int) {
	if o == 0 {
		return
	}
	n := &tr.tr[o]
	n.sum, n.tag = tr.rev(n.sum, n.tag)
	n.rev = !n.rev
}

func (tr *LazyTreap[T, K, M]) pushdown(o int) {
	if o == 0 {
		return
	}
	n := &tr.tr[o]
	tr.apply_tag(n.ls, n.tag)
	tr.apply_tag(n.rs, n.tag)
	n.tag = tr.id()
	if n.rev {
		tr.apply_rev(n.ls)
		tr.apply_rev(n.rs)
		n.rev = false
		n.ls, n.rs = n.rs, n.ls
	}
}

func (tr *LazyTreap[T, K, M]) Split(o, v int) (int, int) {
	if o == 0 {
		return 0, 0
	}
	tr.pushdown(o)
	if v <= tr.sz(tr.ls(o)) {
		lo, ro := tr.Split(tr.ls(o), v)
		tr.tr[o].ls = ro
		tr.pushup(o)
		return lo, o
	} else {
		lo, ro := tr.Split(tr.rs(o), v-tr.sz(tr.ls(o))-1)
		tr.tr[o].rs = lo
		tr.pushup(o)
		return o, ro
	}
}

func (tr *LazyTreap[T, K, M]) Merge(lo, ro int) int {
	if lo == 0 {
		return ro
	}
	if ro == 0 {
		return lo
	}
	if tr.pri(lo) < tr.pri(ro) {
		tr.pushdown(lo)
		tr.tr[lo].rs = tr.Merge(tr.tr[lo].rs, ro)
		tr.pushup(lo)
		return lo
	} else {
		tr.pushdown(ro)
		tr.tr[ro].ls = tr.Merge(lo, tr.tr[ro].ls)
		tr.pushup(ro)
		return ro
	}
}

func (tr *LazyTreap[T, K, M]) Insert(i int, x T) {
	lo, ro := tr.Split(tr.root, i)
	tr.root = tr.Merge(lo, tr.Merge(tr.new_node(x), ro))
}

func (tr *LazyTreap[T, K, M]) Erase(i int) {
	o, ro := tr.Split(tr.root, i+1)
	lo, _ := tr.Split(o, i)
	tr.root = tr.Merge(lo, ro)
}

func (tr *LazyTreap[T, K, M]) Apply(l, r int, tag K) {
	o, ro := tr.Split(tr.root, r)
	lo, m := tr.Split(o, l)
	tr.apply_tag(m, tag)
	tr.root = tr.Merge(lo, tr.Merge(m, ro))
}

func (tr *LazyTreap[T, K, M]) Reverse(l, r int) {
	o, ro := tr.Split(tr.root, r)
	lo, m := tr.Split(o, l)
	tr.apply_rev(m)
	tr.root = tr.Merge(lo, tr.Merge(m, ro))
}

func (tr *LazyTreap[T, K, M]) Prod(l, r int) T {
	o, ro := tr.Split(tr.root, r)
	lo, m := tr.Split(o, l)
	res := tr.sum(m)
	tr.root = tr.Merge(lo, tr.Merge(m, ro))
	return res
}
