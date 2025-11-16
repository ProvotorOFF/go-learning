package product

import (
	"net/http"
	"order-api-start/pkg/req"
	"order-api-start/pkg/res"
	"strconv"
)

type ProductHandler struct {
	repo *ProductRepository
}

type Deps struct {
	Repo *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps Deps) {
	handler := ProductHandler{
		repo: deps.Repo,
	}

	router.HandleFunc("GET /products", handler.list())
	router.HandleFunc("GET /products/{id}", handler.get())
	router.HandleFunc("POST /products", handler.create())
	router.HandleFunc("PUT /products/{id}", handler.update())
	router.HandleFunc("DELETE /products/{id}", handler.delete())
}

func (handler *ProductHandler) get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

		if err != nil {
			res.Json(w, err.Error(), 500)
			return
		}

		product, err := handler.repo.GetById(id)

		if err != nil {
			res.Json(w, err.Error(), 500)
			return
		}

		res.Json(w, product, 200)
	}
}

func (handler *ProductHandler) list() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := handler.repo.All()

		if err != nil {
			res.Json(w, err.Error(), 500)
			return
		}

		res.Json(w, products, 200)
	}
}

func (handler *ProductHandler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		product, err := req.HandleBody[Product](&w, r)
		if err != nil {
			return
		}

		handler.repo.Create(product)
		res.Json(w, product, 201)
	}
}

func (handler *ProductHandler) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

		if err != nil {
			res.Json(w, err.Error(), 400)
			return
		}

		product, err := req.HandleBody[Product](&w, r)
		product.ID = uint(id)

		if err != nil {
			return
		}

		handler.repo.Update(product)
		res.Json(w, product, 200)
	}
}

func (handler *ProductHandler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

		if err != nil {
			res.Json(w, err.Error(), 400)
			return
		}

		err = handler.repo.Delete(id)

		if err != nil {
			res.Json(w, err.Error(), 500)
			return
		}

		res.Json(w, nil, 200)
	}
}
