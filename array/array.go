package array

type ArrayNoConcurrent[T any] struct {
	values []T
}

func NewArrayNoConcurrent[T any](values []T) Array[T] {
	return &ArrayNoConcurrent[T]{
		values,
	}
}

func (a *ArrayNoConcurrent[T]) Map(callback func(T) T) Array[T] {
	for i, v := range a.values {
		value := callback(v)
		a.values[i] = value
	}
	return a
}

func (a *ArrayNoConcurrent[T]) Values() []T {
	return a.values
}
