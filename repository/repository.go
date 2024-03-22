package repository

type repository[T any] interface {
	Create(T) error
	ListAll() ([]T, error)
	List(param string) ([]T, error)
	Update(T) error
	Delete(T) error
	ById(uint) (T, error)
}

type operation[T any] interface {
	repository[T]
	SetStatus() error
	ByField(string) ([]T, error)
}
