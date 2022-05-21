package store

import "github.com/shop/http-rest-api/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}

type ProductRepository interface {
	Create(*model.Product) error
	Find(int) (*model.Product, error)
	List(*model.ProductFilter) ([]*model.Product, error)
}
