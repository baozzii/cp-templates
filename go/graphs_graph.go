package templates

type weight_type interface {
	integer | void
}

type edge[T weight_type] struct {
	w T
	v int32
}

func (e edge[T]) get() (int, T) {
	return int(e.v), e.w
}

type graph[T weight_type] struct {
	__n   int32
	__adj [][]edge[T]
}

func new_graph[T weight_type](n int) *graph[T] {
	return &graph[T]{int32(n), make([][]edge[T], n)}
}

func (g *graph[T]) add_edge(u, v int) {
	var w T
	g.__adj[u] = append(g.__adj[u], edge[T]{w, int32(v)})
}

func (g *graph[T]) from_edges(edges [][]int, offset int) {
	for _, e := range edges {
		g.add_edge(e[0]-offset, e[1]-offset)
	}
}

func (g *graph[T]) add_weighted_edge(u, v int, w T) {
	g.__adj[u] = append(g.__adj[u], edge[T]{w, int32(v)})
}

func (g *graph[T]) from_weighted_edges(edges [][]int, offset int) {
	for _, e := range edges {
		if w, ok := any(e[2]).(T); ok {
			g.add_weighted_edge(e[0]-offset, e[1]-offset, w)
		}
	}
}

func (g *graph[T]) adj(u int) []edge[T] {
	return g.__adj[u]
}

func (g *graph[T]) n() int {
	return int(g.__n)
}

// Only usable after Go1.23
// func (g *graph[T]) adj(u int) iter.Seq2[int, T] {
// 	return func(yield func(int, T) bool) {
// 		for _, e := range g.adj[u] {
// 			if !yield(int(e.v), e.w) {
// 				break
// 			}
// 		}
// 	}
// }
