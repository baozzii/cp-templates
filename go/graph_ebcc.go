package templates

func ebcc_id[T weight_type](g *graph[T]) []int {
	type neighbor struct{ to, eid int32 }
	n := g.n()
	ng := make([][]neighbor, n)
	i := int32(0)
	for u := 0; u < n; u++ {
		for _, e := range g.adj(u) {
			v, _ := e.Get()
			if u < v {
				ng[v] = append(ng[v], neighbor{int32(u), i})
				ng[u] = append(ng[u], neighbor{int32(v), i})
				i++
			}
		}
	}
	b := make([]bool, i)
	dfn := make([]int32, n)
	ts := int32(0)
	var tarjan func(int32, int32) int32
	tarjan = func(u, fid int32) int32 {
		ts++
		dfn[u] = ts
		lu := ts
		for _, e := range ng[u] {
			if v := e.to; dfn[v] == 0 {
				lv := tarjan(v, e.eid)
				lu = min(lu, lv)
				if lv > dfn[u] {
					b[e.eid] = true
				}
			} else if e.eid != fid {
				lu = min(lu, dfn[v])
			}
		}
		return lu
	}
	for v, t := range dfn {
		if t == 0 {
			tarjan(int32(v), -1)
		}
	}
	id := make([]int, n)
	i = 1
	var dfs func(int32)
	dfs = func(u int32) {
		id[u] = int(i)
		for _, e := range ng[u] {
			if !b[e.eid] && id[e.to] == 0 {
				dfs(e.to)
			}
		}
	}
	for j := range id {
		if id[j] == 0 {
			dfs(int32(j))
			i++
		}
	}
	for i := range id {
		id[i]--
	}
	return id
}
