package teststore

import (
	"github.com/shop/http-rest-api/internal/app/model"
	"github.com/shop/http-rest-api/internal/app/store"
)

type ProductRepository struct {
	store    *Store
	products map[int]*model.Product
}

func (r *ProductRepository) Create(p *model.Product) error {
	p.BeforeCreate()
	id := len(r.products) + 1
	r.products[id] = p
	p.ID = id
	return nil
}
func (r *ProductRepository) Find(id int) (*model.Product, error) {
	p, ok := r.products[id]

	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return p, nil
}
func (r *ProductRepository) List(pf *model.ProductFilter) ([]*model.Product, error) {
	res := []*model.Product{}
	for _, v := range r.products {
		if pf != nil && v.Price < pf.MinPrice {
			continue
		}
		if pf != nil && pf.MaxPrice != 0 && v.Price > pf.MaxPrice {
			continue
		}
		if pf != nil && pf.Limit != 0 && len(res) == pf.Limit {
			continue
		}
		res = append(res, v)
	}
	if len(res) == 0 {
		return nil, store.ErrRecordNotFound
	}
	return res, nil

}
