package eroutine

type Result[R any] struct {
	Result R
	Err    error
}

// Routine 协程处理
func Routine[T, R any](ts []T, fnc func(T, chan<- Result[R])) ([]R, error) {
	ch := make(chan Result[R])
	for _, task := range ts {
		go fnc(task, ch)
	}
	results := make([]Result[R], len(ts))
	for i := range results {
		results[i] = <-ch
	}
	var rs []R
	for _, r := range results {
		if r.Err != nil {
			return nil, r.Err
		}
		rs = append(rs, r.Result)
	}
	return rs, nil
}

func NewResult[T any](t T) Result[T] {
	return Result[T]{Result: t}
}

func NewError[T any](err error) Result[T] {
	return Result[T]{Err: err}
}
