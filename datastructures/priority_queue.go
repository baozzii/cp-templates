package datastructures

type PriorityQueueInfo[T any] interface {
	cmp(T, T) int
}

type PriorityQueue[T any, M PriorityQueueInfo[T]] struct {
	hp  []T
	cmp func(T, T) int
}

func NewPriorityQueue[T any, M PriorityQueueInfo[T]](m M) *PriorityQueue[T, M] {
	return &PriorityQueue[T, M]{make([]T, 0), m.cmp}
}

func NewPriorityQueueWith[T any, M PriorityQueueInfo[T]](a []T) *PriorityQueue[T, M] {
	hp := make([]T, len(a))
	copy(hp, a)
	pq := &PriorityQueue[T, M]{hp: hp}
	for i := len(hp)/2 - 1; i >= 0; i-- {
		pq.down(i)
	}
	return pq
}

func (pq *PriorityQueue[T, M]) Empty() bool {
	return len(pq.hp) == 0
}

func (pq *PriorityQueue[T, M]) Top() T {
	return pq.hp[0]
}

func (pq *PriorityQueue[T, M]) Push(x T) {
	pq.hp = append(pq.hp, x)
	pq.up(len(pq.hp) - 1)
}

func (pq *PriorityQueue[T, M]) Pop() {
	n := len(pq.hp)
	pq.hp[0] = pq.hp[n-1]
	pq.hp = pq.hp[:n-1]
	if len(pq.hp) > 0 {
		pq.down(0)
	}
}

func (pq *PriorityQueue[T, M]) less(i, j int) bool {
	return pq.cmp(pq.hp[i], pq.hp[j]) < 0
}

func (pq *PriorityQueue[T, M]) up(i int) {
	for {
		if i == 0 {
			break
		}
		p := (i - 1) / 2
		if !pq.less(i, p) {
			break
		}
		pq.hp[i], pq.hp[p] = pq.hp[p], pq.hp[i]
		i = p
	}
}

func (pq *PriorityQueue[T, M]) down(i int) {
	n := len(pq.hp)
	for {
		left := 2*i + 1
		if left >= n {
			break
		}
		best := left
		right := left + 1
		if right < n && pq.less(right, left) {
			best = right
		}
		if !pq.less(best, i) {
			break
		}
		pq.hp[i], pq.hp[best] = pq.hp[best], pq.hp[i]
		i = best
	}
}
