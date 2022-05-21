package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/shop/http-rest-api/internal/app/store"
)

type Store struct {
	db                *sql.DB
	userRepository    *UserRepository
	productRepository *ProductRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//User....
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

//Product...
func (s *Store) Product() store.ProductRepository {
	if s.productRepository != nil {
		return s.productRepository
	}

	s.productRepository = &ProductRepository{
		store: s,
	}
	return s.productRepository
}
