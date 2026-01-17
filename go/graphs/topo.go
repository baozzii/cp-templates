package graphs

func TopologicalSort[T WeightType](g *Graph[T]) []int {
	deg := make([]int32, g.n)
	for i := range deg {
		for _, e := range g.adj[i] {
			j, _ := e.Get()
			deg[j]++
		}
	}
	q := make([]int32, 0, g.n)
	ans := make([]int, 0, g.n)
	for i := range g.n {
		if deg[i] == 0 {
			q = append(q, int32(i))
		}
	}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		ans = append(ans, int(u))
		for _, e := range g.adj[u] {
			v, _ := e.Get()
			deg[v]--
			if deg[v] == 0 {
				q = append(q, int32(v))
			}
		}
	}
	if len(ans) == g.n {
		return ans
	}
	return []int{}
}
