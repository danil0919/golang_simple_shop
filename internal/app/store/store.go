package store

type Store interface {
	User() UserRepository
	Product() ProductRepository
}
