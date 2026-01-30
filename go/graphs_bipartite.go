package templates

func bipartite[T weight_type](g *graph[T]) ([]int, bool) {
	n := g.n()
	col := make([]int, n)
	vis := make([]bool, n)
	f := true
	c := 0
	for i := range n {
		if vis[i] {
			continue
		}
		var dfs func(int, int)
		dfs = func(u, d int) {
			vis[u] = true
			col[u] = d + 2*c
			for _, e := range g.adj(u) {
				v, _ := e.get()
				if !vis[v] {
					dfs(v, d^1)
				} else {
					if col[u] == col[v] {
						f = false
					}
				}
			}
		}
		dfs(i, 0)
		c++
	}
	return col, f
}
