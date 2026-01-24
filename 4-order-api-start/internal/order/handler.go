package order

import (
	"net/http"
	"order-api-start/configs"
	"order-api-start/pkg/middleware"
	"order-api-start/pkg/req"
	"order-api-start/pkg/res"
	"strconv"
)

type UserRepositoryInterface interface {
	FindIdByPhone(string) (uint, error)
}

type OrderHandler struct {
	repo     *OrderRepository
	conf     *configs.Config
	userRepo UserRepositoryInterface
}

type Deps struct {
	Repo     *OrderRepository
	Conf     *configs.Config
	UserRepo UserRepositoryInterface
}

func NewOrderHandler(router *http.ServeMux, deps Deps) {
	handler := OrderHandler{
		repo:     deps.Repo,
		conf:     deps.Conf,
		userRepo: deps.UserRepo,
	}

	orderMux := http.NewServeMux()

	orderMux.HandleFunc("POST /order", handler.store())
	orderMux.HandleFunc("GET /order/{id}", handler.get())
	orderMux.HandleFunc("GET /my-orders", handler.list())

	router.Handle("/", middleware.Auth(orderMux, handler.conf))
}

func (handler *OrderHandler) store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userPhone, ok := r.Context().Value(middleware.ContextPhoneKey).(string)
		if !ok || userPhone == "" {
			res.Json(w, res.ErrorResponse{Message: "unauthorized"}, http.StatusUnauthorized)
			return
		}

		userId, err := handler.userRepo.FindIdByPhone(userPhone)

		validated, err := req.HandleBody[OrderCreateRequest](&w, r)
		if err != nil {
			return
		}

		order, err := handler.repo.CreateFromRequest(validated, userId)

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
			return
		}

		userPhone, ok := r.Context().Value(middleware.ContextPhoneKey).(string)
		if !ok || userPhone == "" {
			res.Json(w, res.ErrorResponse{Message: "unauthorized"}, http.StatusUnauthorized)
			return
		}

		userId, err := handler.userRepo.FindIdByPhone(userPhone)

		if err != nil {
			res.Json(w, res.ErrorResponse{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		order, err := handler.repo.GetById(id)

		if userId != order.UserID {
			res.Json(w, res.ErrorResponse{Message: "Forbidden"}, http.StatusForbidden)
		}

		res.Json(w, order, http.StatusOK)
	}
}

func (handler *OrderHandler) list() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userPhone, ok := r.Context().Value(middleware.ContextPhoneKey).(string)
		if !ok || userPhone == "" {
			res.Json(w, res.ErrorResponse{Message: "unauthorized"}, http.StatusUnauthorized)
			return
		}

		userId, err := handler.userRepo.FindIdByPhone(userPhone)

		if err != nil {
			res.Json(w, res.ErrorResponse{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		orders, err := handler.repo.FindForUser(userId)

		if err != nil {
			res.Json(w, res.ErrorResponse{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		res.Json(w, orders, http.StatusOK)
	}
}
