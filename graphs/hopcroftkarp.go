package graphs

type HopcroftKarp struct {
	nl, nr int32
	ed     [][2]int32
	ml, mr []int32
}

func NewHopcroftKarp(nl, nr int) *HopcroftKarp {
	ml := make([]int32, nl)
	mr := make([]int32, nr)
	for i := range ml {
		ml[i] = -1
	}
	for i := range mr {
		mr[i] = -1
	}
	return &HopcroftKarp{int32(nl), int32(nr), make([][2]int32, 0), ml, mr}
}

func (h *HopcroftKarp) AddEdge(u, v int) {
	h.ed = append(h.ed, [2]int32{int32(u), int32(v)})
}

func (h *HopcroftKarp) Matching() int {
	xorshift := func(x uint32) uint32 {
		x ^= x << 13
		x ^= x >> 17
		x ^= x << 5
		return x
	}
	x := uint32(114514)
	for i := range h.ed {
		j := x % uint32(i+1)
		h.ed[i], h.ed[j] = h.ed[j], h.ed[i]
		x = xorshift(x)
	}
	n := h.nl
	g := make([]int32, len(h.ed))
	deg := make([]int32, n+1)
	for _, e := range h.ed {
		deg[e[0]]++
	}
	for i := int32(0); i < n; i++ {
		deg[i+1] += deg[i]
	}
	for _, e := range h.ed {
		deg[e[0]]--
		g[deg[e[0]]] = e[1]
	}
	a := make([]int32, n)
	p := make([]int32, n)
	q := make([]int32, n)
	for {
		for i := int32(0); i < n; i++ {
			a[i] = -1
			p[i] = -1
		}
		t := 0
		for i := int32(0); i < n; i++ {
			if h.ml[i] == -1 {
				q[t] = i
				t++
				a[i] = i
				p[i] = i
			}
		}

		f := false
		for i := 0; i < t; i++ {
			x := q[i]
			if h.ml[a[x]] != -1 {
				continue
			}
			for j := deg[x]; j < deg[x+1]; j++ {
				y := g[j]
				if h.mr[y] == -1 {
					for y != -1 {
						h.mr[y] = x
						y, h.ml[x] = h.ml[x], y
						x = p[x]
					}
					f = true
					break
				}
				nx := h.mr[y]
				if p[nx] == -1 {
					q[t] = nx
					t++
					p[nx] = x
					a[nx] = a[x]
				}
			}
		}
		if !f {
			break
		}
	}
	cnt := 0
	for i := range h.ml {
		if h.ml[i] != -1 {
			cnt++
		}
	}
	return cnt
}

func (h *HopcroftKarp) MatchLeft(u int) int {
	return int(h.ml[u])
}

func (h *HopcroftKarp) MatchRight(u int) int {
	return int(h.mr[u])
}
