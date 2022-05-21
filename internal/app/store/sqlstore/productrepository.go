package sqlstore

import (
	"database/sql"
	"fmt"

	"github.com/shop/http-rest-api/internal/app/model"
	"github.com/shop/http-rest-api/internal/app/store"
)

type ProductRepository struct {
	store *Store
}

func (r *ProductRepository) Create(p *model.Product) error {
	p.BeforeCreate()
	return r.store.db.QueryRow(
		"INSERT INTO products (created,name,price,description) VALUES ($1,$2,$3,$4) RETURNING id",
		p.Created,
		p.Name,
		p.Price,
		p.Description,
	).Scan(&p.ID)
}
func (r *ProductRepository) Find(id int) (*model.Product, error) {
	p := &model.Product{}
	if err := r.store.db.QueryRow(
		"SELECT id, created, name, description, price FROM products WHERE id = $1",
		id,
	).Scan(
		&p.ID,
		&p.Created,
		&p.Name,
		&p.Description,
		&p.Price,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return p, nil
}
func (r *ProductRepository) List(pf *model.ProductFilter) ([]*model.Product, error) {
	where := ""
	limit := ""
	if pf != nil {

		where += fmt.Sprintf(" WHERE price > %v ", pf.MinPrice)
		if pf.MaxPrice != 0 {
			where += fmt.Sprintf(" AND price < %v ", pf.MaxPrice)
		}
		if pf.Limit != 0 {
			limit += fmt.Sprintf(" limit %d ", pf.Limit)
		}

	}

	rows, err := r.store.db.Query("SELECT id, created, name, description, price FROM products" + where + limit)
	if err != nil {
		return nil, err
	}
	result := []*model.Product{}
	for rows.Next() {
		product := &model.Product{}
		if err := rows.Scan(
			&product.ID,
			&product.Created,
			&product.Name,
			&product.Description,
			&product.Price,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}
			return nil, err
		}

		result = append(result, product)
	}

	return result, nil
}
