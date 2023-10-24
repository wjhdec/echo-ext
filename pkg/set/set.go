package set

type Set[T comparable] interface {
	Add(item ...T)
	Contains(item T) bool
	Delete(item T)
}

type set[T comparable] map[T]any

// New 新建set
func New[T comparable](items ...T) Set[T] {
	s := &set[T]{}
	for _, item := range items {
		s.Add(item)
	}
	return s
}

func NewWithSlice[T comparable](items []T) Set[T] {
	s := &set[T]{}
	for _, item := range items {
		s.Add(item)
	}
	return s
}

func (s set[T]) Add(v ...T) {
	for _, e := range v {
		s[e] = nil
	}
}

func (s set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s set[T]) Delete(v T) {
	if s.Contains(v) {
		delete(s, v)
	}
}
