package datastructures

type BiTrieNode struct {
	son [2]int
	cnt int
}

type BiTrie struct {
	t []BiTrieNode
	b int
}

func NewBiTrie(b int) *BiTrie {
	t := make([]BiTrieNode, 2)
	return &BiTrie{t, b}
}

func (t *BiTrie) new_node() int {
	t.t = append(t.t, BiTrieNode{})
	return len(t.t) - 1
}

func (t *BiTrie) Insert(x int) int {
	u := 1
	t.t[u].cnt++
	for i := t.b - 1; i >= 0; i-- {
		v := x >> i & 1
		if t.t[u].son[v] == 0 {
			t.t[u].son[v] = t.new_node()
		}
		u = t.t[u].son[v]
		t.t[u].cnt++
	}
	return u
}

func (t *BiTrie) Erase(x int) {
	u := 1
	t.t[u].cnt--
	for i := t.b - 1; i >= 0; i-- {
		v := x >> i & 1
		u = t.t[u].son[v]
		t.t[u].cnt--
	}
}

func (t *BiTrie) MaxXor(x int) int {
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

func (t *BiTrie) MinXor(x int) int {
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
