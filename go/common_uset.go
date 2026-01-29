package templates

type uset[T comparable] map[T]struct{}

func (s *uset[T]) insert(x T) {
	(*s)[x] = struct{}{}
}

func (s *uset[T]) erase(x T) {
	delete(*s, x)
}

func (s *uset[T]) contains(x T) bool {
	_, ok := (*s)[x]
	return ok
}

func (s *uset[T]) size() int {
	return len(*s)
}

func (s *uset[T]) empty() bool {
	return s.size() == 0
}

func (s *uset[T]) clear() {
	clear(*s)
}

func (s *uset[T]) keys() vector[T] {
	b := make(vector[T], 0, s.size())
	for v := range *s {
		b.push_back(v)
	}
	return b
}

func (s *uset[T]) copy() uset[T] {
	t := make(uset[T])
	t.union(*s)
	return t
}

func (s *uset[T]) union(t uset[T]) {
	for v := range t {
		s.insert(v)
	}
}

func (s *uset[T]) intersect(t uset[T]) {
	ns := make(uset[T])
	for v := range t {
		if s.contains(v) {
			ns.insert(v)
		}
	}
	*s = ns
}

func union[T comparable](s, t uset[T]) uset[T] {
	ns := s.copy()
	ns.union(t)
	return ns
}

func intersect[T comparable](s, t uset[T]) uset[T] {
	ns := make(uset[T])
	for v := range s {
		if t.contains(v) {
			ns.insert(v)
		}
	}
	return ns
}
