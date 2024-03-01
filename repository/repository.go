package repository

type repository[T any] interface {
	Create(T) error
	ListAll() ([]T, error)
	List(param string) ([]T, error)
	Update(T) error
	Delete(T) error
}
