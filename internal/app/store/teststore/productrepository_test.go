package teststore_test

import (
	"testing"

	"github.com/shop/http-rest-api/internal/app/model"
	"github.com/shop/http-rest-api/internal/app/store"
	"github.com/shop/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestProductRepository_Create(t *testing.T) {
	s := teststore.New()
	p := model.TestProduct(t)
	err := s.Product().Create(p)

	assert.NoError(t, err)
	assert.NotNil(t, p)
}

func TestProductRespository_Find(t *testing.T) {
	s := teststore.New()
	p := model.TestProduct(t)
	s.Product().Create(p)

	res, err := s.Product().Find(p.ID)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	res, err = s.Product().Find(p.ID + 1)
	assert.Error(t, store.ErrRecordNotFound)
	assert.Nil(t, res)

}

func TestProductRepository_List(t *testing.T) {
	s := teststore.New()
	p := model.TestProduct(t)
	s.Product().Create(p)
	p.Price = 6000
	s.Product().Create(p)

	res, err := s.Product().List(nil)
	assert.NotEmpty(t, res)
	assert.NoError(t, err)

	res, err = s.Product().List(&model.ProductFilter{MinPrice: 4000})
	t.Log(res)

	assert.NoError(t, err)
	assert.Equal(t, p.ID, res[0].ID)

	res, err = s.Product().List(&model.ProductFilter{Limit: 1})
	assert.NoError(t, err)
	assert.Equal(t, len(res), 1)

}
