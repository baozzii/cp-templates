package graphs

import . "cp-templates/go/common"

func SccId[T Integer | Void](g *Graph[T]) []int {
	n := int32(g.n)
	t := int32(0)
	sid := int32(0)
	vis := make([]int32, 0, n)
	low := make([]int32, n)
	ord := make([]int32, n)
	ids := make([]int, n)
	for i := range ord {
		ord[i] = -1
	}
	var dfs func(int32)
	dfs = func(u int32) {
		low[u] = t
		ord[u] = t
		t++
		vis = append(vis, u)
		for _, v := range g.adj[u] {
			if ord[v.v] == -1 {
				dfs(int32(v.v))
				low[u] = min(low[u], low[v.v])
			} else {
				low[u] = min(low[u], ord[v.v])
			}
		}
		if low[u] == ord[u] {
			for {
				v := vis[len(vis)-1]
				vis = vis[:len(vis)-1]
				ord[v] = n
				ids[v] = int(sid)
				if u == v {
					break
				}
			}
			sid++
		}
	}
	for i := range n {
		if ord[i] == -1 {
			dfs(i)
		}
	}
	for i := range ids {
		ids[i] = int(sid) - 1 - ids[i]
	}
	return ids
}
