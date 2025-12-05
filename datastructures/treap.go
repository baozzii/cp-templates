package datastructures

import "math/rand"

type TreapInfo[T any] interface {
	op(T, T) T
	e() T
}

type TreapNode[T any, M TreapInfo[T]] struct {
	ls, rs, sz int
	pri        uint
	val, sum   T
}

type Treap[T any, M TreapInfo[T]] struct {
	root int
	tr   []TreapNode[T, M]
	op   func(T, T) T
	e    func() T
}

func NewTreap[T any, M TreapInfo[T]](m M) *Treap[T, M] {
	tr := make([]TreapNode[T, M], 1)
	tr[0] = TreapNode[T, M]{0, 0, 0, 0, m.e(), m.e()}
	return &Treap[T, M]{0, tr, m.op, m.e}
}

func (tr *Treap[T, M]) new_node(x T) int {
	tr.tr = append(tr.tr, TreapNode[T, M]{0, 0, 1, uint(rand.Uint64()), x, x})
	return len(tr.tr) - 1
}

func (tr *Treap[T, M]) ls(o int) int {
	return tr.tr[o].ls
}

func (tr *Treap[T, M]) rs(o int) int {
	return tr.tr[o].rs
}

func (tr *Treap[T, M]) sz(o int) int {
	return tr.tr[o].sz
}

func (tr *Treap[T, M]) val(o int) T {
	return tr.tr[o].val
}

func (tr *Treap[T, M]) sum(o int) T {
	return tr.tr[o].sum
}

func (tr *Treap[T, M]) pri(o int) uint {
	return tr.tr[o].pri
}

func (tr *Treap[T, M]) pushup(o int) {
	tr.tr[o].sz = 1 + tr.sz(tr.ls(o)) + tr.sz(tr.rs(o))
	tr.tr[o].sum = tr.op(tr.sum(tr.ls(o)), tr.op(tr.val(o), tr.sum(tr.rs(o))))
}

func (tr *Treap[T, M]) Split(o, v int) (int, int) {
	if o == 0 {
		return 0, 0
	}
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

func (tr *Treap[T, M]) Merge(lo, ro int) int {
	if lo == 0 {
		return ro
	}
	if ro == 0 {
		return lo
	}
	if tr.pri(lo) < tr.pri(ro) {
		tr.tr[lo].rs = tr.Merge(tr.tr[lo].rs, ro)
		tr.pushup(lo)
		return lo
	} else {
		tr.tr[ro].ls = tr.Merge(lo, tr.tr[ro].ls)
		tr.pushup(ro)
		return ro
	}
}

func (tr *Treap[T, M]) Insert(i int, x T) {
	lo, ro := tr.Split(tr.root, i)
	tr.root = tr.Merge(lo, tr.Merge(tr.new_node(x), ro))
}

func (tr *Treap[T, M]) Erase(i int, x T) {
	o, ro := tr.Split(tr.root, i+1)
	lo, _ := tr.Split(o, i)
	tr.root = tr.Merge(lo, ro)
}

func (tr *Treap[T, M]) Prod(l, r int) T {
	o, ro := tr.Split(tr.root, r)
	lo, m := tr.Split(o, l)
	res := tr.sum(m)
	tr.root = tr.Merge(lo, tr.Merge(m, ro))
	return res
}
