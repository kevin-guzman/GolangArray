package array

type Array[T any] interface {
	// NewArray(values []T) Array[T]
	Map(callback func(T) T) Array[T]
	Values() []T
	// ForEach(callback func(T))
	// Filter(callback func(T) bool) *Array[T]
	// Concat(arrays ...[]T) *Array[T]
	// Push(values ...T) *Array[T]
}
