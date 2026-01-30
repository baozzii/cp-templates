package templates

type twosat struct {
	n int
	g *graph[void]
}

func new_twosat(n int) *twosat {
	return &twosat{n, new_graph[void](2 * n)}
}

func (t *twosat) add_clause(i int, f bool, j int, g bool) {
	t.g.add_edge(2*i+cond(f, 0, 1), 2*j+cond(g, 1, 0))
	t.g.add_edge(2*j+cond(g, 0, 1), 2*i+cond(f, 1, 0))
}

func (t *twosat) satisfiable() bool {
	id := scc_id(t.g)
	for i := range t.n {
		if id[i*2] == id[i*2+1] {
			return false
		}
	}
	return true
}

func (t *twosat) answer() []bool {
	id := scc_id(t.g)
	ans := make([]bool, t.n)
	for i := range t.n {
		ans[i] = id[i*2] < id[i*2+1]
	}
	return ans
}
