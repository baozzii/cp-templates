package graphs

import . "cp-templates/common"

type Edge[T Integer | Void] struct {
	w T
	v int
}

type Graph[T Integer | Void] struct {
	n   int
	adj [][]Edge[T]
}

func NewGraph[T Integer | Void](n int) *Graph[T] {
	return &Graph[T]{n, make([][]Edge[T], n)}
}

func (g *Graph[T]) AddEdge(u, v int) {
	var w T
	g.adj[u] = append(g.adj[u], Edge[T]{w, v})
}

func (g *Graph[T]) AddWeightedEdge(u, v int, w T) {
	g.adj[u] = append(g.adj[u], Edge[T]{w, v})
}
