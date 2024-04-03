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
	Create(T) error
	ListAll() ([]T, error)
	Update(uint) error
	Delete(uint) error
	ByAgent(uint) ([]T, error)
	ByProduct(uint) ([]T, error)
}
