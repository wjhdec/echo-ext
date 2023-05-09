package pair

type Pair[T, U any] struct {
	Key   T
	Value U
}

func New[T, U any](key T, value U) *Pair[T, U] {
	return &Pair[T, U]{key, value}
}
