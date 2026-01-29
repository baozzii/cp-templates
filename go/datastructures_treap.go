package templates

import "math/rand"

type treap_info[T any] interface {
	op(T, T) T
	e() T
}

type treap_node[T any, M treap_info[T]] struct {
	ls, rs, sz int
	pri        uint
	val, sum   T
}

type treap[T any, M treap_info[T]] struct {
	root int
	tr   []treap_node[T, M]
	op   func(T, T) T
	e    func() T
}

func new_treap[T any, M treap_info[T]](m M) *treap[T, M] {
	tr := make([]treap_node[T, M], 1)
	tr[0] = treap_node[T, M]{0, 0, 0, 0, m.e(), m.e()}
	return &treap[T, M]{0, tr, m.op, m.e}
}

func (tr *treap[T, M]) __new_node(x T) int {
	tr.tr = append(tr.tr, treap_node[T, M]{0, 0, 1, uint(rand.Uint64()), x, x})
	return len(tr.tr) - 1
}

func (tr *treap[T, M]) ls(o int) int {
	return tr.tr[o].ls
}

func (tr *treap[T, M]) rs(o int) int {
	return tr.tr[o].rs
}

func (tr *treap[T, M]) sz(o int) int {
	return tr.tr[o].sz
}

func (tr *treap[T, M]) val(o int) T {
	return tr.tr[o].val
}

func (tr *treap[T, M]) sum(o int) T {
	return tr.tr[o].sum
}

func (tr *treap[T, M]) pri(o int) uint {
	return tr.tr[o].pri
}

func (tr *treap[T, M]) __pushup(o int) {
	tr.tr[o].sz = 1 + tr.sz(tr.ls(o)) + tr.sz(tr.rs(o))
	tr.tr[o].sum = tr.op(tr.sum(tr.ls(o)), tr.op(tr.val(o), tr.sum(tr.rs(o))))
}

func (tr *treap[T, M]) split(o, v int) (int, int) {
	if o == 0 {
		return 0, 0
	}
	if v <= tr.sz(tr.ls(o)) {
		lo, ro := tr.split(tr.ls(o), v)
		tr.tr[o].ls = ro
		tr.__pushup(o)
		return lo, o
	} else {
		lo, ro := tr.split(tr.rs(o), v-tr.sz(tr.ls(o))-1)
		tr.tr[o].rs = lo
		tr.__pushup(o)
		return o, ro
	}
}

func (tr *treap[T, M]) merge(lo, ro int) int {
	if lo == 0 {
		return ro
	}
	if ro == 0 {
		return lo
	}
	if tr.pri(lo) < tr.pri(ro) {
		tr.tr[lo].rs = tr.merge(tr.tr[lo].rs, ro)
		tr.__pushup(lo)
		return lo
	} else {
		tr.tr[ro].ls = tr.merge(lo, tr.tr[ro].ls)
		tr.__pushup(ro)
		return ro
	}
}

func (tr *treap[T, M]) insert(i int, x T) {
	lo, ro := tr.split(tr.root, i)
	tr.root = tr.merge(lo, tr.merge(tr.__new_node(x), ro))
}

func (tr *treap[T, M]) erase(i int) {
	o, ro := tr.split(tr.root, i+1)
	lo, _ := tr.split(o, i)
	tr.root = tr.merge(lo, ro)
}

func (tr *treap[T, M]) prod(l, r int) T {
	o, ro := tr.split(tr.root, r)
	lo, m := tr.split(o, l)
	res := tr.sum(m)
	tr.root = tr.merge(lo, tr.merge(m, ro))
	return res
}
