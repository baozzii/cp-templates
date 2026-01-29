package templates

type lichao_line[K integer, T integer] struct{ k, b T }

func (ln lichao_line[K, T]) Eval(x K) T { return ln.k*T(x) + ln.b }

type lichao_segtree[K integer, T integer] struct {
	L, R  K
	isMin bool
	def   lichao_line[K, T]

	lines []lichao_line[K, T]
	lc    []int32
	rc    []int32
}

func new_lichao_segtree[K integer, T integer](L, R K, isMin bool) *lichao_segtree[K, T] {
	t := &lichao_segtree[K, T]{L: L, R: R, isMin: isMin}
	if isMin {
		t.def = lichao_line[K, T]{0, limit[T]().max()}
	} else {
		t.def = lichao_line[K, T]{0, limit[T]().min()}
	}
	t.lines = []lichao_line[K, T]{t.def}
	t.lc = []int32{-1}
	t.rc = []int32{-1}
	return t
}

func (t *lichao_segtree[K, T]) __new_node() int32 {
	t.lines = append(t.lines, t.def)
	t.lc = append(t.lc, -1)
	t.rc = append(t.rc, -1)
	return int32(len(t.lines) - 1)
}

func (t *lichao_segtree[K, T]) insert(f lichao_line[K, T]) { t.insert_seg(t.L, t.R, f) }

func (t *lichao_segtree[K, T]) insert_seg(a, b K, f lichao_line[K, T]) {
	if b < t.L || t.R < a {
		return
	}
	a = max(a, t.L)
	b = min(b, t.R)
	if a > b {
		return
	}
	cmp := func(x, y T) bool {
		if t.isMin {
			return x < y
		}
		return x > y
	}
	var addLine func(int32, K, K, lichao_line[K, T])
	addLine = func(v int32, l, r K, g lichao_line[K, T]) {
		m := l + (r-l)/2
		if cmp(g.Eval(m), t.lines[v].Eval(m)) {
			t.lines[v], g = g, t.lines[v]
		}
		if l == r {
			return
		}
		if !cmp(g.Eval(l), t.lines[v].Eval(l)) && !cmp(g.Eval(r), t.lines[v].Eval(r)) {
			return
		}
		if cmp(g.Eval(l), t.lines[v].Eval(l)) {
			if t.lc[v] == -1 {
				t.lc[v] = t.__new_node()
			}
			addLine(t.lc[v], l, m, g)
		} else {
			if t.rc[v] == -1 {
				t.rc[v] = t.__new_node()
			}
			addLine(t.rc[v], m+1, r, g)
		}
	}

	var dfs func(int32, K, K)
	dfs = func(v int32, l, r K) {
		if r < a || b < l {
			return
		}
		if a <= l && r <= b {
			addLine(v, l, r, f)
			return
		}
		if l == r {
			if cmp(f.Eval(l), t.lines[v].Eval(l)) {
				t.lines[v] = f
			}
			return
		}
		m := l + (r-l)/2
		if t.lc[v] == -1 {
			t.lc[v] = t.__new_node()
		}
		if t.rc[v] == -1 {
			t.rc[v] = t.__new_node()
		}
		dfs(t.lc[v], l, m)
		dfs(t.rc[v], m+1, r)
	}
	dfs(0, t.L, t.R)
}

func (t *lichao_segtree[K, T]) query(x K) T {
	cmp := func(a, b T) bool {
		if t.isMin {
			return a < b
		}
		return a > b
	}
	var dfs func(int32, K, K) T
	dfs = func(v int32, l, r K) T {
		if v == -1 {
			return t.def.Eval(x)
		}
		res := t.lines[v].Eval(x)
		if l == r {
			return res
		}
		m := l + (r-l)/2
		sub := cond(x <= m, dfs(t.lc[v], l, m), dfs(t.rc[v], m+1, r))
		if cmp(sub, res) {
			return sub
		}
		return res
	}
	return dfs(0, t.L, t.R)
}
