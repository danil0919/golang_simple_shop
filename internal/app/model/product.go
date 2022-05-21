package model

import "time"

type Product struct {
	ID          int       `json:"id"`
	Created     time.Time `json:"-"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
}

type ProductFilter struct {
	MinPrice float64 `json:"min_price"`
	MaxPrice float64 `json:"max_price"`
	Limit    int     `json:"limit"`
}

func (p *Product) BeforeCreate() {
	p.Created = time.Now()
}
