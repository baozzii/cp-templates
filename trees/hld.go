package trees

type HLD struct {
	t                          *Tree
	in, out, hvy, top, sz, ord []int
	dfn                        int
}

func NewHLD(t *Tree) *HLD {
	h := &HLD{
		t:   t,
		in:  make([]int, t.n),
		out: make([]int, t.n),
		hvy: make([]int, t.n),
		top: make([]int, t.n),
		sz:  make([]int, t.n),
		ord: make([]int, t.n),
	}
	for i := 0; i < t.n; i++ {
		h.hvy[i] = -1
	}
	h.build()
	return h
}

func (h *HLD) build() {
	var dfs1 func(int)
	dfs1 = func(u int) {
		h.sz[u] = 1
		for _, v := range h.t.adj[u] {
			h.t.pa[v] = u
			h.t.dep[v] = h.t.dep[u] + 1
			dfs1(v)
			h.sz[u] += h.sz[v]
			if h.hvy[u] == -1 || h.sz[v] > h.sz[h.hvy[u]] {
				h.hvy[u] = v
			}
		}
	}
	var dfs2 func(int, int)
	dfs2 = func(u, tp int) {
		h.top[u] = tp
		h.in[u] = h.dfn
		h.ord[h.dfn] = u
		h.dfn++
		if h.hvy[u] != -1 {
			dfs2(h.hvy[u], tp)
		}
		for _, v := range h.t.adj[u] {
			if v != h.hvy[u] {
				dfs2(v, v)
			}
		}
		h.out[u] = h.dfn
	}
	dfs1(h.t.root)
	h.dfn = 0
	dfs2(h.t.root, h.t.root)
}

func (h *HLD) Subtree(u int) (l, r int) {
	return h.in[u], h.out[u]
}

func (h *HLD) Lca(u, v int) int {
	for h.top[u] != h.top[v] {
		if h.t.dep[h.top[u]] > h.t.dep[h.top[v]] {
			u = h.t.pa[h.top[u]]
		} else {
			v = h.t.pa[h.top[v]]
		}
	}
	if h.t.dep[u] < h.t.dep[v] {
		return u
	}
	return v
}

func (h *HLD) DoForPathDirected(u, v int, lca bool, f func(l, r int, rev bool)) {
	type seg struct {
		l, r int
		rev  bool
	}
	down := make([]seg, 0)
	for h.top[u] != h.top[v] {
		if h.t.dep[h.top[u]] >= h.t.dep[h.top[v]] {
			top := h.top[u]
			l := h.in[top]
			r := h.in[u] + 1
			f(l, r, true)
			u = h.t.pa[top]
		} else {
			top := h.top[v]
			l := h.in[top]
			r := h.in[v] + 1
			down = append(down, seg{l, r, false})
			v = h.t.pa[top]
		}
	}
	if h.t.dep[u] >= h.t.dep[v] {
		l := h.in[v]
		if !lca {
			l++
		}
		r := h.in[u] + 1
		if l < r {
			f(l, r, true)
		}
	} else {
		l := h.in[u]
		if !lca {
			l++
		}
		r := h.in[v] + 1
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

func (h *HLD) DoForPath(u, v int, lca bool, f func(l, r int)) {
	h.DoForPathDirected(u, v, lca, func(l, r int, rev bool) { f(l, r) })
}

func (h *HLD) DoForSubtree(u int, f func(int, int)) {
	l, r := h.Subtree(u)
	f(l, r)
}

func (h *HLD) KthPa(u, k int) int {
	if k < 0 {
		return -1
	}
	for u != -1 {
		top := h.top[u]
		dist := h.t.dep[u] - h.t.dep[top]
		if k <= dist {
			return h.ord[h.in[u]-k]
		}
		k -= dist + 1
		u = h.t.pa[top]
	}
	return -1
}

func (h *HLD) MoveTo(u, v, k int) int {
	if k == 0 {
		return u
	}
	l := h.Lca(u, v)
	du := h.t.dep[u] - h.t.dep[l]
	dv := h.t.dep[v] - h.t.dep[l]
	if k >= du+dv {
		return v
	}
	if k <= du {
		return h.KthPa(u, k)
	}
	return h.KthPa(v, du+dv-k)
}

func (h *HLD) IsAnc(u, v int) bool {
	return h.in[u] <= h.in[v] && h.out[u] >= h.out[v]
}
