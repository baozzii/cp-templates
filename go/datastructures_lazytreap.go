package templates

import "math/rand"

type lazytreap_info[T any, K any] interface {
	op(T, T) T
	e() T
	mp(K, T) T
	cp(K, K) K
	id() K
	rev(T, K) (T, K)
}

type lazytreap_node[T any, K any, M lazytreap_info[T, K]] struct {
	ls, rs, sz int
	pri        uint
	val, sum   T
	tag        K
	rev        bool
}

type lazytreap[T any, K any, M lazytreap_info[T, K]] struct {
	root int
	tr   []lazytreap_node[T, K, M]
	op   func(T, T) T
	e    func() T
	mp   func(K, T) T
	cp   func(K, K) K
	id   func() K
	rev  func(T, K) (T, K)
}

func new_lazytreap[T any, K any, M lazytreap_info[T, K]](m M) *lazytreap[T, K, M] {
	tr := make([]lazytreap_node[T, K, M], 1)
	tr[0] = lazytreap_node[T, K, M]{0, 0, 0, 0, m.e(), m.e(), m.id(), false}
	return &lazytreap[T, K, M]{0, tr, m.op, m.e, m.mp, m.cp, m.id, m.rev}
}

func (tr *lazytreap[T, K, M]) __new_node(x T) int {
	tr.tr = append(tr.tr, lazytreap_node[T, K, M]{0, 0, 1, uint(rand.Uint64()), x, x, tr.id(), false})
	return len(tr.tr) - 1
}

func (tr *lazytreap[T, K, M]) ls(o int) int {
	return tr.tr[o].ls
}

func (tr *lazytreap[T, K, M]) rs(o int) int {
	return tr.tr[o].rs
}

func (tr *lazytreap[T, K, M]) sz(o int) int {
	return tr.tr[o].sz
}

func (tr *lazytreap[T, K, M]) val(o int) T {
	return tr.tr[o].val
}

func (tr *lazytreap[T, K, M]) sum(o int) T {
	return tr.tr[o].sum
}

func (tr *lazytreap[T, K, M]) pri(o int) uint {
	return tr.tr[o].pri
}

func (tr *lazytreap[T, K, M]) __pushup(o int) {
	tr.tr[o].sz = 1 + tr.sz(tr.ls(o)) + tr.sz(tr.rs(o))
	tr.tr[o].sum = tr.op(tr.sum(tr.ls(o)), tr.op(tr.val(o), tr.sum(tr.rs(o))))
}

func (tr *lazytreap[T, K, M]) __apply_tag(o int, tag K) {
	if o == 0 {
		return
	}
	n := &tr.tr[o]
	n.sum = tr.mp(tag, n.sum)
	n.val = tr.mp(tag, n.val)
	n.tag = tr.cp(tag, n.tag)
}

func (tr *lazytreap[T, K, M]) __apply_rev(o int) {
	if o == 0 {
		return
	}
	n := &tr.tr[o]
	n.sum, n.tag = tr.rev(n.sum, n.tag)
	n.rev = !n.rev
}

func (tr *lazytreap[T, K, M]) __pushdown(o int) {
	if o == 0 {
		return
	}
	n := &tr.tr[o]
	tr.__apply_tag(n.ls, n.tag)
	tr.__apply_tag(n.rs, n.tag)
	n.tag = tr.id()
	if n.rev {
		tr.__apply_rev(n.ls)
		tr.__apply_rev(n.rs)
		n.rev = false
		n.ls, n.rs = n.rs, n.ls
	}
}

func (tr *lazytreap[T, K, M]) split(o, v int) (int, int) {
	if o == 0 {
		return 0, 0
	}
	tr.__pushdown(o)
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

func (tr *lazytreap[T, K, M]) merge(lo, ro int) int {
	if lo == 0 {
		return ro
	}
	if ro == 0 {
		return lo
	}
	if tr.pri(lo) < tr.pri(ro) {
		tr.__pushdown(lo)
		tr.tr[lo].rs = tr.merge(tr.tr[lo].rs, ro)
		tr.__pushup(lo)
		return lo
	} else {
		tr.__pushdown(ro)
		tr.tr[ro].ls = tr.merge(lo, tr.tr[ro].ls)
		tr.__pushup(ro)
		return ro
	}
}

func (tr *lazytreap[T, K, M]) insert(i int, x T) {
	lo, ro := tr.split(tr.root, i)
	tr.root = tr.merge(lo, tr.merge(tr.__new_node(x), ro))
}

func (tr *lazytreap[T, K, M]) erase(i int) {
	o, ro := tr.split(tr.root, i+1)
	lo, _ := tr.split(o, i)
	tr.root = tr.merge(lo, ro)
}

func (tr *lazytreap[T, K, M]) apply(l, r int, tag K) {
	o, ro := tr.split(tr.root, r)
	lo, m := tr.split(o, l)
	tr.__apply_tag(m, tag)
	tr.root = tr.merge(lo, tr.merge(m, ro))
}

func (tr *lazytreap[T, K, M]) reverse(l, r int) {
	o, ro := tr.split(tr.root, r)
	lo, m := tr.split(o, l)
	tr.__apply_rev(m)
	tr.root = tr.merge(lo, tr.merge(m, ro))
}

func (tr *lazytreap[T, K, M]) prod(l, r int) T {
	o, ro := tr.split(tr.root, r)
	lo, m := tr.split(o, l)
	res := tr.sum(m)
	tr.root = tr.merge(lo, tr.merge(m, ro))
	return res
}
