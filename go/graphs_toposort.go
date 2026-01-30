package templates

func toposort[T weight_type](g *graph[T]) ([]int, bool) {
	n := g.n()
	deg := make([]int32, n)
	for i := range deg {
		for _, e := range g.adj(i) {
			j, _ := e.Get()
			deg[j]++
		}
	}
	q := make([]int32, 0, n)
	ans := make([]int, 0, n)
	for i := range n {
		if deg[i] == 0 {
			q = append(q, int32(i))
		}
	}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		ans = append(ans, int(u))
		for _, e := range g.adj(int(u)) {
			v, _ := e.get()
			deg[v]--
			if deg[v] == 0 {
				q = append(q, int32(v))
			}
		}
	}
	return ans, len(ans) == n
}
