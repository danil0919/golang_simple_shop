package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/shop/http-rest-api/internal/app/model"
	"github.com/shop/http-rest-api/internal/app/store"
)

func (s *server) handleProductsList() http.HandlerFunc {
	type request struct {
		MinPrice float64 `schema:"min_price"`
		MaxPrice float64 `schema:"max_price"`
		Limit    int     `schema:"limit"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		pf := model.ProductFilter{}
		decoder := schema.NewDecoder()
		request := &request{}
		err := decoder.Decode(request, r.URL.Query())
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		pf = model.ProductFilter(*request)

		products, err := s.store.Product().List(&pf)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
		if len(products) == 0 {
			s.error(w, r, http.StatusNotFound, store.ErrRecordNotFound)
			return
		}

		s.respond(w, r, http.StatusOK, products)
	}
}
func (s *server) handleProductsFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		p, err := s.store.Product().Find(id)
		if err != nil {
			if err == store.ErrRecordNotFound {
				s.error(w, r, http.StatusNotFound, err)
			}
			s.error(w, r, http.StatusInternalServerError, err)
		}
		s.respond(w, r, http.StatusOK, p)
	}
}

func (s *server) handleProductsCreate() http.HandlerFunc {
	type request struct {
		Name        string  `json:"name"`
		Price       float64 `json:"price"`
		Description string  `json:"description"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		product := &model.Product{
			Name:        req.Name,
			Price:       req.Price,
			Description: req.Description,
		}

		if err := s.store.Product().Create(product); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, product)
	}
}
