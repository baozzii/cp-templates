package common

type Umap[T comparable, K any] map[T]K

func (s *Umap[T, K]) Erase(x T) {
	delete(*s, x)
}

func (s *Umap[T, K]) Contains(x T) bool {
	_, ok := (*s)[x]
	return ok
}

func (s *Umap[T, K]) Size() int {
	return len(*s)
}

func (s *Umap[T, K]) Empty() bool {
	return s.Size() == 0
}

func (s *Umap[T, K]) Clear() {
	clear(*s)
}

func (s *Umap[T, K]) Keys() *Vec[T] {
	b := make(Vec[T], 0, s.Size())
	for v := range *s {
		b.PushBack(v)
	}
	return &b
}

func (s *Umap[T, K]) Values() *Vec[K] {
	b := make(Vec[K], 0, s.Size())
	for _, v := range *s {
		b.PushBack(v)
	}
	return &b
}
