package templates

type umap[T comparable, K any] map[T]K

func (s *umap[T, K]) erase(x T) {
	delete(*s, x)
}

func (s *umap[T, K]) contains(x T) bool {
	_, ok := (*s)[x]
	return ok
}

func (s *umap[T, K]) size() int {
	return len(*s)
}

func (s *umap[T, K]) empty() bool {
	return s.size() == 0
}

func (s *umap[T, K]) clear() {
	clear(*s)
}

func (s *umap[T, K]) keys() vector[T] {
	b := make(vector[T], 0, s.size())
	for v := range *s {
		b.push_back(v)
	}
	return b
}

func (s *umap[T, K]) values() vector[K] {
	b := make(vector[K], 0, s.size())
	for _, v := range *s {
		b.push_back(v)
	}
	return b
}
