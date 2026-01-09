package graphs

import (
	. "cp-templates/go/common"
)

type WeightType interface {
	Integer | Void
}

type Edge[T WeightType] struct {
	w T
	v int32
}

func (e Edge[T]) Get() (int, T) {
	return int(e.v), e.w
}

type Graph[T WeightType] struct {
	n   int
	adj [][]Edge[T]
}

func NewGraph[T WeightType](n int) *Graph[T] {
	return &Graph[T]{n, make([][]Edge[T], n)}
}

func (g *Graph[T]) AddEdge(u, v int) {
	var w T
	g.adj[u] = append(g.adj[u], Edge[T]{w, int32(v)})
}

func (g *Graph[T]) FromEdges(edges [][]int, offset int) {
	for _, e := range edges {
		g.AddEdge(e[0]-offset, e[1]-offset)
	}
}

func (g *Graph[T]) AddWeightedEdge(u, v int, w T) {
	g.adj[u] = append(g.adj[u], Edge[T]{w, int32(v)})
}

func (g *Graph[T]) FromWeightedEdges(edges [][]int, offset int) {
	for _, e := range edges {
		if w, ok := any(e[2]).(T); ok {
			g.AddWeightedEdge(e[0]-offset, e[1]-offset, w)
		}
	}
}

// Only usable after Go1.23
// func (g *Graph[T]) Adj(u int) iter.Seq2[int, T] {
// 	return func(yield func(int, T) bool) {
// 		for _, e := range g.adj[u] {
// 			if !yield(int(e.v), e.w) {
// 				break
// 			}
// 		}
// 	}
// }
