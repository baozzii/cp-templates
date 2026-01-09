package graphs

import (
	"container/heap"
	. "cp-templates/go/common"
)

type DijkstraPair[T Integer] struct {
	d T
	v int32
}

type DijkstraHeap[T Integer] []DijkstraPair[T]

func (h DijkstraHeap[T]) Len() int           { return len(h) }
func (h DijkstraHeap[T]) Less(i, j int) bool { return h[i].d < h[j].d }
func (h DijkstraHeap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *DijkstraHeap[T]) Push(x any)        { *h = append(*h, x.(DijkstraPair[T])) }
func (h *DijkstraHeap[T]) Pop() any          { a := *h; x := a[len(a)-1]; *h = a[:len(a)-1]; return x }

func Dijkstra[T Integer](g *Graph[T], s int) []T {
	inf := Limit[T]().Max() / 4
	dis := make([]T, g.n)
	for i := range dis {
		dis[i] = inf
	}
	dis[s] = 0
	h := make(DijkstraHeap[T], 0)
	heap.Push(&h, DijkstraPair[T]{0, int32(s)})
	for h.Len() > 0 {
		cur := heap.Pop(&h).(DijkstraPair[T])
		if cur.d != dis[cur.v] {
			continue
		}
		u := cur.v
		for _, e := range g.adj[u] {
			v, w := e.Get()
			nd := cur.d + w
			if nd < dis[v] {
				dis[v] = nd
				heap.Push(&h, DijkstraPair[T]{nd, int32(v)})
			}
		}
	}
	return dis
}
