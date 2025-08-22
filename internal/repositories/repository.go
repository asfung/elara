package repositories

type Repository[T any] interface {
	Create(entity T) (T, error)
	Update(entity T) (T, error)
	FindById(id any) (*T, error)
	Delete(id any) error
}
