package templates

type avl_info[T any] interface {
	cmp(T, T) int
}

type avl_node[T any, M avl_info[T]] struct {
	key  T
	cnt  int32
	h, s int32
	l, r int32
}

type avl[T any, M avl_info[T]] struct {
	t    []avl_node[T, M]
	cmp  func(T, T) int
	root int32
}

func new_avl[T any, M avl_info[T]](m M) *avl[T, M] {
	t := make([]avl_node[T, M], 1)
	return &avl[T, M]{t, m.cmp, 0}
}

func (t *avl[T, M]) __new_node(x T, c int) int32 {
	t.t = append(t.t, avl_node[T, M]{x, int32(c), 1, int32(c), 0, 0})
	return int32(len(t.t) - 1)
}

func (t *avl[T, M]) __pushup(o int32) {
	if o != 0 {
		t.t[o].h = max(t.t[t.t[o].l].h, t.t[t.t[o].r].h) + 1
		t.t[o].s = t.t[t.t[o].l].s + t.t[t.t[o].r].s + t.t[o].cnt
	}
}

func (t *avl[T, M]) __balance(o int32) int32 {
	return t.t[t.t[o].l].h - t.t[t.t[o].r].h
}

func (t *avl[T, M]) __rotate_right(y int32) int32 {
	x := t.t[y].l
	z := t.t[x].r
	t.t[x].r = y
	t.t[y].l = z
	t.__pushup(y)
	t.__pushup(x)
	return x
}

func (t *avl[T, M]) __rotate_left(y int32) int32 {
	x := t.t[y].r
	z := t.t[x].l
	t.t[x].l = y
	t.t[y].r = z
	t.__pushup(y)
	t.__pushup(x)
	return x
}

func (t *avl[T, M]) rebalance(o int32) int32 {
	if o != 0 {
		t.__pushup(o)
		b := t.__balance(o)
		if b > 1 {
			if t.__balance(t.t[o].l) < 0 {
				t.t[o].l = t.__rotate_left(t.t[o].l)
			}
			return t.__rotate_right(o)
		}
		if b < -1 {
			if t.__balance(t.t[o].r) > 0 {
				t.t[o].r = t.__rotate_right(t.t[o].r)
			}
			return t.__rotate_left(o)
		}
	}
	return o
}

func (t *avl[T, M]) insert(x T) {
	var insert func(int32, T) int32
	insert = func(o int32, x T) int32 {
		if o == 0 {
			return t.__new_node(x, 1)
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

func (t *avl[T, M]) erase(x T) {
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

func (t *avl[T, M]) count_lt(x T) int {
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

func (t *avl[T, M]) count_le(x T) int {
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

func (t *avl[T, M]) count_gt(x T) int {
	return t.size() - t.count_le(x)
}

func (t *avl[T, M]) count_ge(x T) int {
	return t.size() - t.count_lt(x)
}

func (t *avl[T, M]) CountBetween(x, y T) int {
	return t.count_gt(y) - t.count_lt(x)
}

func (t *avl[T, M]) get(k int) T {
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

func (t *avl[T, M]) count(x T) int {
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

func (t *avl[T, M]) find_lt(x, d T) T {
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

func (t *avl[T, M]) find_le(x, d T) T {
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

func (t *avl[T, M]) find_gt(x, d T) T {
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

func (t *avl[T, M]) find_ge(x, d T) T {
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

func (t *avl[T, M]) size() int {
	return int(t.t[t.root].s)
}

func (t *avl[T, M]) min() T {
	o := t.root
	for {
		if t.t[o].l != 0 {
			o = t.t[o].l
		} else {
			return t.t[o].key
		}
	}
}

func (t *avl[T, M]) max() T {
	o := t.root
	for {
		if t.t[o].r != 0 {
			o = t.t[o].r
		} else {
			return t.t[o].key
		}
	}
}
