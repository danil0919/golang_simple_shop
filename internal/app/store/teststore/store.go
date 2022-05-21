package teststore

import (
	_ "github.com/lib/pq"
	"github.com/shop/http-rest-api/internal/app/model"
	"github.com/shop/http-rest-api/internal/app/store"
)

type Store struct {
	userRepository    *UserRepository
	productRepository *ProductRepository
}

func New() *Store {
	return &Store{}
}

//User....
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}

//User....
func (s *Store) Product() store.ProductRepository {
	if s.productRepository != nil {
		return s.productRepository
	}

	s.productRepository = &ProductRepository{
		store:    s,
		products: make(map[int]*model.Product),
	}

	return s.productRepository
}
