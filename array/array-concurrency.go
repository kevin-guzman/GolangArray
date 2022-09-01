package array

import "sync"

type ArrayConcurrent[T any] struct {
	values []T
}

func NewArrayConcurrent[T any](values []T) Array[T] {
	return &ArrayConcurrent[T]{
		values: values,
	}
}

func (a *ArrayConcurrent[T]) Map(callback func(T) T) Array[T] {
	var wg sync.WaitGroup
	for i, v := range a.values {
		wg.Add(1)
		go func() {
			defer wg.Done()
			value := callback(v)
			a.values[i] = value
		}()
	}
	wg.Wait()
	return a
}

func (a *ArrayConcurrent[T]) Values() []T {
	return a.values
}
