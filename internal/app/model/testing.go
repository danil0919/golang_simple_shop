package model

import (
	"testing"
	"time"
)

//TestUser...
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "Password",
	}
}

func TestProduct(t *testing.T) *Product {
	return &Product{
		Name:        "Samsung s10",
		Description: "Cool device",
		Price:       3000,
		Created:     time.Now(),
	}
}
