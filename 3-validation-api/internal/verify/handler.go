package verify

import (
	"net/http"
	"validation-api/configs"
)

type verifyHandler struct {
	conf *configs.Config
}

func NewVerifyHandler(router *http.ServeMux, conf *configs.Config) {
	handler := verifyHandler{
		conf: conf,
	}
	router.HandleFunc("POST /send", handler.send())
	router.HandleFunc("GET  /verify/{hash}", handler.verify())
}

func (handler *verifyHandler) send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}

func (handler *verifyHandler) verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
