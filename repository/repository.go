package repository

type repository[T any] interface {
	Create(T) error
	ListAll() ([]T, error)
}
