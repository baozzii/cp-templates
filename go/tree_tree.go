package templates

const (
	TREE_DFN = 1 << iota
	TREE_SZ
	TREE_DEP
	TREE_HLD
)

type tree struct {
	__n, __root                      int32
	__pa                             []int32
	__in, __out, __ord, __hvy, __top []int32
	__dep                            []int32
	__sz                             []int32
	__adj                            [][]int
}

func (t *tree) n() int {
	return int(t.__n)
}

func (t *tree) root() int {
	return int(t.__root)
}

func (t *tree) pa(u int) int {
	return int(t.__pa[u])
}

func (t *tree) in(u int) int {
	return int(t.__in[u])
}

func (t *tree) out(u int) int {
	return int(t.__out[u])
}

func (t *tree) ord(u int) int {
	return int(t.__ord[u])
}

func (t *tree) hvy(u int) int {
	return int(t.__hvy[u])
}

func (t *tree) top(u int) int {
	return int(t.__top[u])
}

func (t *tree) dep(u int) int {
	return int(t.__dep[u])
}

func (t *tree) sz(u int) int {
	return int(t.__sz[u])
}

func (t *tree) adj(u int) []int {
	return t.__adj[u]
}

func new_tree(n int) *tree {
	return &tree{__n: int32(n), __adj: make([][]int, n), __pa: make([]int32, n)}
}

func (t *tree) add_edge(u, v int) {
	t.__adj[u] = append(t.__adj[u], v)
	t.__adj[v] = append(t.__adj[v], u)
}

func (t *tree) set_root(u int) {
	t.__root = int32(u)
}

func (t *tree) build(f int) {
	n := t.n()
	contains := func(x int) bool {
		return f&x != 0
	}
	if contains(TREE_HLD) {
		t.__dep = make([]int32, n)
		t.__sz = make([]int32, n)
		for i := range t.__sz {
			t.__sz[i] = 1
		}
		t.__in = make([]int32, n)
		t.__out = make([]int32, n)
		t.__ord = make([]int32, n)
		t.__hvy = make([]int32, n)
		for i := range t.__hvy {
			t.__hvy[i] = -1
		}
		t.__top = make([]int32, n)
		var dfs1 func(int, int)
		dfs1 = func(u, p int) {
			t.__pa[u] = int32(p)
			for _, v := range t.adj(u) {
				if v == p {
					continue
				}
				t.__dep[v] = t.__dep[u] + 1
				dfs1(v, u)
				t.__sz[u] += t.__sz[v]
				if t.__hvy[u] == -1 || t.__sz[v] > t.__sz[t.__hvy[u]] {
					t.__hvy[u] = int32(v)
				}
			}
		}
		ts := int32(0)
		var dfs2 func(int32, int32)
		dfs2 = func(u, tp int32) {
			t.__top[u] = tp
			t.__in[u] = ts
			t.__ord[ts] = u
			ts++
			if t.__hvy[u] != -1 {
				dfs2(t.__hvy[u], tp)
			}
			for _, v := range t.adj(int(u)) {
				if int32(v) == t.__pa[u] {
					continue
				}
				if int32(v) != t.__hvy[u] {
					dfs2(int32(v), int32(v))
				}
			}
			t.__out[u] = ts
		}
		dfs1(t.root(), -1)
		dfs2(int32(t.root()), int32(t.root()))
	} else {
		if contains(TREE_DEP) {
			t.__dep = make([]int32, n)
		}
		if contains(TREE_SZ) {
			t.__sz = make([]int32, n)
			for i := range t.__sz {
				t.__sz[i] = 1
			}
		}
		if contains(TREE_DFN) {
			t.__in = make([]int32, n)
			t.__out = make([]int32, n)
			t.__ord = make([]int32, n)
		}
		ts := int32(0)
		var dfs func(int, int)
		dfs = func(u, p int) {
			t.__pa[u] = int32(p)
			if contains(TREE_DFN) {
				t.__in[u] = ts
				t.__ord[ts] = int32(u)
				ts++
			}
			for _, v := range t.adj(u) {
				if v == p {
					continue
				}
				if contains(TREE_DEP) {
					t.__dep[v] = t.__dep[u] + 1
				}
				dfs(v, u)
				if contains(TREE_SZ) {
					t.__sz[u] += t.__sz[v]
				}
			}
			if contains(TREE_DFN) {
				t.__out[u] = ts
			}
		}
		dfs(t.root(), -1)
	}
	for u := range n {
		for i, v := range t.__adj[u] {
			if int32(v) == t.__pa[u] {
				t.__adj[u] = append(t.__adj[u][:i], t.__adj[u][i+1:]...)
				break
			}
		}
	}
}

func (t *tree) lca(u, v int) int {
	for t.top(u) != t.top(v) {
		if t.dep(t.top(u)) > t.dep(t.top(v)) {
			u = t.pa(t.top(u))
		} else {
			v = t.pa(t.top(v))
		}
	}
	if t.dep(u) < t.dep(v) {
		return u
	}
	return v
}

func (t *tree) do_for_path_directed(u, v int, lca bool, f func(l, r int, rev bool)) {
	type seg struct {
		l, r int
		rev  bool
	}
	down := make([]seg, 0)
	for t.top(u) != t.top(v) {
		if t.dep(t.top(u)) >= t.dep(t.top(v)) {
			top := t.top(u)
			l := t.in(top)
			r := t.in(u) + 1
			f(l, r, true)
			u = t.pa(top)
		} else {
			top := t.top(v)
			l := t.in(top)
			r := t.in(v) + 1
			down = append(down, seg{l, r, false})
			v = t.pa(top)
		}
	}
	if t.dep(u) >= t.dep(v) {
		l := t.in(v)
		if !lca {
			l++
		}
		r := t.in(u) + 1
		if l < r {
			f(l, r, true)
		}
	} else {
		l := t.in(u)
		if !lca {
			l++
		}
		r := t.in(v) + 1
		if l < r {
			down = append(down, seg{l, r, false})
		}
	}
	for i := len(down) - 1; i >= 0; i-- {
		s := down[i]
		if s.l < s.r {
			f(s.l, s.r, s.rev)
		}
	}
}

func (t *tree) do_for_path(u, v int, lca bool, f func(l, r int)) {
	t.do_for_path_directed(u, v, lca, func(l, r int, rev bool) { f(l, r) })
}

func (t *tree) kth_pa(u, k int) int {
	if k < 0 {
		return -1
	}
	for u != -1 {
		top := t.top(u)
		dist := t.dep(u) - t.dep(top)
		if k <= dist {
			return t.ord(t.in(u) - k)
		}
		k -= dist + 1
		u = t.pa(top)
	}
	return -1
}

func (t *tree) move_to(u, v, k int) int {
	if k == 0 {
		return u
	}
	l := t.lca(u, v)
	du := t.dep(u) - t.dep(l)
	dv := t.dep(v) - t.dep(l)
	if k >= du+dv {
		return v
	}
	if k <= du {
		return t.kth_pa(u, k)
	}
	return t.kth_pa(v, du+dv-k)
}
