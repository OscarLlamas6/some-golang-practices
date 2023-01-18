package utils

type Iterator[T any] interface {
	Next() bool
	Value() T
}
