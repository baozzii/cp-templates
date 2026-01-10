package datastructures

type AVLInfo[T any] interface {
	cmp(T, T) int
}

type AVLNode[T any, M AVLInfo[T]] struct {
	key  T
	cnt  int32
	h, s int32
	l, r int32
}

type AVL[T any, M AVLInfo[T]] struct {
	t    []AVLNode[T, M]
	cmp  func(T, T) int
	root int32
}

func NewAVL[T any, M AVLInfo[T]](m M) *AVL[T, M] {
	t := make([]AVLNode[T, M], 1)
	return &AVL[T, M]{t, m.cmp, 0}
}

func (t *AVL[T, M]) new_node(x T, c int) int32 {
	t.t = append(t.t, AVLNode[T, M]{x, int32(c), 1, int32(c), 0, 0})
	return int32(len(t.t) - 1)
}

func (t *AVL[T, M]) pushup(o int32) {
	if o != 0 {
		t.t[o].h = max(t.t[t.t[o].l].h, t.t[t.t[o].r].h) + 1
		t.t[o].s = t.t[t.t[o].l].s + t.t[t.t[o].r].s + t.t[o].cnt
	}
}

func (t *AVL[T, M]) balance(o int32) int32 {
	return t.t[t.t[o].l].h - t.t[t.t[o].r].h
}

func (t *AVL[T, M]) rotate_right(y int32) int32 {
	x := t.t[y].l
	z := t.t[x].r
	t.t[x].r = y
	t.t[y].l = z
	t.pushup(y)
	t.pushup(x)
	return x
}

func (t *AVL[T, M]) rotate_left(y int32) int32 {
	x := t.t[y].r
	z := t.t[x].l
	t.t[x].l = y
	t.t[y].r = z
	t.pushup(y)
	t.pushup(x)
	return x
}

func (t *AVL[T, M]) rebalance(o int32) int32 {
	if o != 0 {
		t.pushup(o)
		b := t.balance(o)
		if b > 1 {
			if t.balance(t.t[o].l) < 0 {
				t.t[o].l = t.rotate_left(t.t[o].l)
			}
			return t.rotate_right(o)
		}
		if b < -1 {
			if t.balance(t.t[o].r) > 0 {
				t.t[o].r = t.rotate_right(t.t[o].r)
			}
			return t.rotate_left(o)
		}
	}
	return o
}

func (t *AVL[T, M]) Insert(x T) {
	var insert func(int32, T) int32
	insert = func(o int32, x T) int32 {
		if o == 0 {
			return t.new_node(x, 1)
		}
		p := t.cmp(x, t.t[o].key)
		if p < 0 {
			t.t[o].l = insert(t.t[o].l, x)
		} else if p > 0 {
			t.t[o].r = insert(t.t[o].r, x)
		} else {
			t.t[o].cnt++
		}
		return t.rebalance(o)
	}
	t.root = insert(t.root, x)
}

func (t *AVL[T, M]) Erase(x T) {
	var nk T
	var nc int32
	var rm func(int32) int32
	rm = func(o int32) int32 {
		if t.t[o].l == 0 {
			nk = t.t[o].key
			nc = t.t[o].cnt
			return t.t[o].r
		}
		t.t[o].l = rm(t.t[o].l)
		return t.rebalance(o)
	}

	var erase func(int32, T) int32
	erase = func(o int32, x T) int32 {
		if o == 0 {
			return 0
		}
		p := t.cmp(x, t.t[o].key)
		if p < 0 {
			t.t[o].l = erase(t.t[o].l, x)
		} else if p > 0 {
			t.t[o].r = erase(t.t[o].r, x)
		} else {
			if t.t[o].cnt > 1 {
				t.t[o].cnt--
			} else {
				if t.t[o].l == 0 || t.t[o].r == 0 {
					return max(t.t[o].l, t.t[o].r)
				} else {
					t.t[o].r = rm(t.t[o].r)
					t.t[o].key = nk
					t.t[o].cnt = nc
				}
			}
		}
		return t.rebalance(o)
	}
	t.root = erase(t.root, x)
}

func (t *AVL[T, M]) CountLt(x T) int {
	res := 0
	o := t.root
	for o != 0 {
		p := t.cmp(x, t.t[o].key)
		if p <= 0 {
			o = t.t[o].l
		} else {
			res += int(t.t[t.t[o].l].s + t.t[o].cnt)
			o = t.t[o].r
		}
	}
	return res
}

func (t *AVL[T, M]) CountLe(x T) int {
	res := 0
	o := t.root
	for o != 0 {
		p := t.cmp(x, t.t[o].key)
		if p < 0 {
			o = t.t[o].l
		} else {
			res += int(t.t[t.t[o].l].s + t.t[o].cnt)
			o = t.t[o].r
		}
	}
	return res
}

func (t *AVL[T, M]) CountGt(x T) int {
	return t.Size() - t.CountLe(x)
}

func (t *AVL[T, M]) CountGe(x T) int {
	return t.Size() - t.CountLt(x)
}

func (t *AVL[T, M]) CountBetween(x, y T) int {
	return t.CountLt(y) - t.CountLt(x)
}

func (t *AVL[T, M]) Get(k int) T {
	o := t.root
	for o != 0 {
		left := t.t[t.t[o].l].s
		if k < int(left) {
			o = t.t[o].l
		} else if k < int(left+t.t[o].cnt) {
			return t.t[o].key
		} else {
			k -= int(left + t.t[o].cnt)
			o = t.t[o].r
		}
	}
	panic(-1)
}

func (t *AVL[T, M]) Count(x T) int {
	o := t.root
	for o != 0 {
		p := t.cmp(x, t.t[o].key)
		if p == 0 {
			return int(t.t[o].cnt)
		} else if p < 0 {
			o = t.t[o].l
		} else {
			o = t.t[o].r
		}
	}
	return 0
}

func (t *AVL[T, M]) FindLt(x, d T) T {
	o := t.root
	ans := 0
	for o != 0 {
		p := t.cmp(x, t.t[o].key)
		if p <= 0 {
			o = t.t[o].l
		} else {
			ans = int(o)
			o = t.t[o].r
		}
	}
	if ans == 0 {
		return d
	}
	return t.t[ans].key
}

func (t *AVL[T, M]) FindLe(x, d T) T {
	o := t.root
	ans := 0
	for o != 0 {
		p := t.cmp(x, t.t[o].key)
		if p < 0 {
			o = t.t[o].l
		} else {
			ans = int(o)
			o = t.t[o].r
		}
	}
	if ans == 0 {
		return d
	}
	return t.t[ans].key
}

func (t *AVL[T, M]) FindGt(x, d T) T {
	o := t.root
	ans := 0
	for o != 0 {
		p := t.cmp(x, t.t[o].key)
		if p < 0 {
			ans = int(o)
			o = t.t[o].l
		} else {
			o = t.t[o].r
		}
	}
	if ans == 0 {
		return d
	}
	return t.t[ans].key
}

func (t *AVL[T, M]) FindGe(x, d T) T {
	o := t.root
	ans := 0
	for o != 0 {
		p := t.cmp(x, t.t[o].key)
		if p <= 0 {
			ans = int(o)
			o = t.t[o].l
		} else {
			o = t.t[o].r
		}
	}
	if ans == 0 {
		return d
	}
	return t.t[ans].key
}

func (t *AVL[T, M]) Size() int {
	return int(t.t[t.root].s)
}

func (t *AVL[T, M]) Min() T {
	o := t.root
	for {
		if t.t[o].l != 0 {
			o = t.t[o].l
		} else {
			return t.t[o].key
		}
	}
}

func (t *AVL[T, M]) Max() T {
	o := t.root
	for {
		if t.t[o].r != 0 {
			o = t.t[o].r
		} else {
			return t.t[o].key
		}
	}
}

// Only usable after Go1.23
// func (t *AVL[T, M]) Values() iter.Seq[T] {
// 	return func(yield func(T) bool) {
// 		var dfs func(int32) bool
// 		dfs = func(u int32) bool {
// 			if u == 0 {
// 				return true
// 			}
// 			if !dfs(t.t[u].l) {
// 				return false
// 			}
// 			for range t.t[u].cnt {
// 				if !yield(t.t[u].key) {
// 					return false
// 				}
// 			}
// 			if !dfs(t.t[u].r) {
// 				return false
// 			}
// 			return true
// 		}
// 		dfs(t.root)
// 	}
// }
