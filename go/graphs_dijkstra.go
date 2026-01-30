package templates

import "container/heap"

type __dij_pair[T integer] struct {
	d T
	v int32
}

type __dij_heap[T integer] []__dij_pair[T]

func (h __dij_heap[T]) Len() int           { return len(h) }
func (h __dij_heap[T]) Less(i, j int) bool { return h[i].d < h[j].d }
func (h __dij_heap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *__dij_heap[T]) Push(x any)        { *h = append(*h, x.(__dij_pair[T])) }
func (h *__dij_heap[T]) Pop() any          { a := *h; x := a[len(a)-1]; *h = a[:len(a)-1]; return x }

func dijkstra[T integer](g *graph[T], s int) []T {
	inf := limit[T]().max() / 4
	dis := make([]T, g.n())
	for i := range dis {
		dis[i] = inf
	}
	dis[s] = 0
	h := make(__dij_heap[T], 0)
	heap.Push(&h, __dij_pair[T]{0, int32(s)})
	for h.Len() > 0 {
		cur := heap.Pop(&h).(__dij_pair[T])
		if cur.d != dis[cur.v] {
			continue
		}
		u := cur.v
		for _, e := range g.adj(int(u)) {
			v, w := e.get()
			nd := cur.d + w
			if nd < dis[v] {
				dis[v] = nd
				heap.Push(&h, __dij_pair[T]{nd, int32(v)})
			}
		}
	}
	return dis
}
