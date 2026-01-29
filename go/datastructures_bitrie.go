package templates

type bitrie_node struct {
	son [2]int
	cnt int
}

type bitrie struct {
	t []bitrie_node
	b int
}

func Newbitrie(b int) *bitrie {
	t := make([]bitrie_node, 2)
	return &bitrie{t, b}
}

func (t *bitrie) __new_node() int {
	t.t = append(t.t, bitrie_node{})
	return len(t.t) - 1
}

func (t *bitrie) insert(x int) int {
	u := 1
	t.t[u].cnt++
	for i := t.b - 1; i >= 0; i-- {
		v := x >> i & 1
		if t.t[u].son[v] == 0 {
			t.t[u].son[v] = t.__new_node()
		}
		u = t.t[u].son[v]
		t.t[u].cnt++
	}
	return u
}

func (t *bitrie) erase(x int) {
	u := 1
	t.t[u].cnt--
	for i := t.b - 1; i >= 0; i-- {
		v := x >> i & 1
		u = t.t[u].son[v]
		t.t[u].cnt--
	}
}

func (t *bitrie) max_xor(x int) int {
	if len(t.t) == 2 {
		return x
	}
	y := 0
	u := 1
	for i := t.b - 1; i >= 0; i-- {
		v := x >> i & 1
		if t.t[u].son[v^1] != 0 && t.t[t.t[u].son[v^1]].cnt > 0 {
			y |= 1 << i
			u = t.t[u].son[v^1]
		} else {
			u = t.t[u].son[v]
		}
	}
	return y
}

func (t *bitrie) min_xor(x int) int {
	if len(t.t) == 2 {
		return x
	}
	y := 0
	u := 1
	for i := t.b - 1; i >= 0; i-- {
		v := x >> i & 1
		if t.t[u].son[v] != 0 && t.t[t.t[u].son[v]].cnt > 0 {
			u = t.t[u].son[v]
		} else {
			y |= 1 << i
			u = t.t[u].son[v^1]
		}
	}
	return y
}
