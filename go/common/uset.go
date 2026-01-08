package common

type Uset[T comparable] map[T]struct{}

func (s *Uset[T]) Insert(x T) {
	(*s)[x] = struct{}{}
}

func (s *Uset[T]) Erase(x T) {
	delete(*s, x)
}

func (s *Uset[T]) Contains(x T) bool {
	_, ok := (*s)[x]
	return ok
}

func (s *Uset[T]) Size() int {
	return len(*s)
}

func (s *Uset[T]) Empty() bool {
	return s.Size() == 0
}

func (s *Uset[T]) Clear() {
	clear(*s)
}

func (s *Uset[T]) Keys() Vec[T] {
	b := make(Vec[T], 0, s.Size())
	for v := range *s {
		b.PushBack(v)
	}
	return b
}

func (s *Uset[T]) FromVec(v Vec[T]) {
	for _, w := range v {
		s.Insert(w)
	}
}
