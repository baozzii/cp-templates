package templates

type priority_queue_info[T any] interface {
	cmp(T, T) int
}

type priority_queue[T any, M priority_queue_info[T]] struct {
	hp  []T
	cmp func(T, T) int
}

func new_priority_queue[T any, M priority_queue_info[T]](m M) *priority_queue[T, M] {
	return &priority_queue[T, M]{make([]T, 0), m.cmp}
}

func new_priority_queue_with[T any, M priority_queue_info[T], E ~[]T](a E) *priority_queue[T, M] {
	hp := make([]T, len(a))
	copy(hp, a)
	pq := &priority_queue[T, M]{hp: hp}
	for i := len(hp)/2 - 1; i >= 0; i-- {
		pq.__down(i)
	}
	return pq
}

func (pq *priority_queue[T, M]) empty() bool {
	return len(pq.hp) == 0
}

func (pq *priority_queue[T, M]) top() T {
	return pq.hp[0]
}

func (pq *priority_queue[T, M]) push(x T) {
	pq.hp = append(pq.hp, x)
	pq.__up(len(pq.hp) - 1)
}

func (pq *priority_queue[T, M]) pop() {
	n := len(pq.hp)
	pq.hp[0] = pq.hp[n-1]
	pq.hp = pq.hp[:n-1]
	if len(pq.hp) > 0 {
		pq.__down(0)
	}
}

func (pq *priority_queue[T, M]) __less(i, j int) bool {
	return pq.cmp(pq.hp[i], pq.hp[j]) < 0
}

func (pq *priority_queue[T, M]) __up(i int) {
	for {
		if i == 0 {
			break
		}
		p := (i - 1) / 2
		if !pq.__less(i, p) {
			break
		}
		pq.hp[i], pq.hp[p] = pq.hp[p], pq.hp[i]
		i = p
	}
}

func (pq *priority_queue[T, M]) __down(i int) {
	n := len(pq.hp)
	for {
		left := 2*i + 1
		if left >= n {
			break
		}
		best := left
		right := left + 1
		if right < n && pq.__less(right, left) {
			best = right
		}
		if !pq.__less(best, i) {
			break
		}
		pq.hp[i], pq.hp[best] = pq.hp[best], pq.hp[i]
		i = best
	}
}
