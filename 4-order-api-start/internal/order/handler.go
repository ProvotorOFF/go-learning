package order

import (
	"net/http"
	"order-api-start/configs"
	"order-api-start/pkg/middleware"
	"order-api-start/pkg/req"
	"order-api-start/pkg/res"
	"strconv"
)

type OrderHandler struct {
	repo *OrderRepository
	conf *configs.Config
}

type Deps struct {
	Repo *OrderRepository
	Conf *configs.Config
}

func NewOrderHandler(router *http.ServeMux, deps Deps) {
	handler := OrderHandler{
		repo: deps.Repo,
		conf: deps.Conf,
	}

	orderMux := http.NewServeMux()

	orderMux.HandleFunc("POST /order", handler.store())
	orderMux.HandleFunc("GET /order/{id}", handler.get())
	orderMux.HandleFunc("GET /my-orders", handler.list())

	router.Handle("/", middleware.Auth(orderMux, handler.conf))
}

func (handler *OrderHandler) store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validated, err := req.HandleBody[OrderCreateRequest](&w, r)
		if err != nil {
			return
		}

		order, err := handler.repo.CreateFromRequest(validated)

		if err != nil {
			res.Json(w, res.ErrorResponse{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		res.Json(w, order, http.StatusCreated)
	}
}

func (handler *OrderHandler) get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

		if err != nil {
			res.Json(w, res.ErrorResponse{Message: err.Error()}, http.StatusBadRequest)
		}

		userPhone, ok := r.Context().Value(middleware.ContextPhoneKey).(string)
		if !ok || userPhone == "" {
			return
		}

		order, err := handler.repo.GetById(id)

		res.Json(w, order, http.StatusOK)
	}
}

func (handler *OrderHandler) list() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
