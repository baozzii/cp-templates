package datastructures

import . "cp-templates/go/common"

type LichaoLine[K Integer, T Integer] struct{ k, b T }

func (ln LichaoLine[K, T]) Eval(x K) T { return ln.k*T(x) + ln.b }

type LiChaoSeg[K Integer, T Integer] struct {
	L, R  K
	isMin bool
	def   LichaoLine[K, T]

	lines []LichaoLine[K, T]
	lc    []int32
	rc    []int32
}

func NewLiChaoSeg[K Integer, T Integer](L, R K, isMin bool) *LiChaoSeg[K, T] {
	t := &LiChaoSeg[K, T]{L: L, R: R, isMin: isMin}
	if isMin {
		t.def = LichaoLine[K, T]{0, Limit[T]().Max()}
	} else {
		t.def = LichaoLine[K, T]{0, Limit[T]().Min()}
	}
	t.lines = []LichaoLine[K, T]{t.def}
	t.lc = []int32{-1}
	t.rc = []int32{-1}
	return t
}

func (t *LiChaoSeg[K, T]) new_node() int32 {
	t.lines = append(t.lines, t.def)
	t.lc = append(t.lc, -1)
	t.rc = append(t.rc, -1)
	return int32(len(t.lines) - 1)
}

func (t *LiChaoSeg[K, T]) Insert(f LichaoLine[K, T]) { t.InsertSeg(t.L, t.R, f) }

func (t *LiChaoSeg[K, T]) InsertSeg(a, b K, f LichaoLine[K, T]) {
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
	var addLine func(int32, K, K, LichaoLine[K, T])
	addLine = func(v int32, l, r K, g LichaoLine[K, T]) {
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
				t.lc[v] = t.new_node()
			}
			addLine(t.lc[v], l, m, g)
		} else {
			if t.rc[v] == -1 {
				t.rc[v] = t.new_node()
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
			t.lc[v] = t.new_node()
		}
		if t.rc[v] == -1 {
			t.rc[v] = t.new_node()
		}
		dfs(t.lc[v], l, m)
		dfs(t.rc[v], m+1, r)
	}
	dfs(0, t.L, t.R)
}

func (t *LiChaoSeg[K, T]) Query(x K) T {
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
		sub := Cond(x <= m, dfs(t.lc[v], l, m), dfs(t.rc[v], m+1, r))
		if cmp(sub, res) {
			return sub
		}
		return res
	}
	return dfs(0, t.L, t.R)
}
