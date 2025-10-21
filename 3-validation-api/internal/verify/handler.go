package verify

import (
	"fmt"
	"net/http"
	"validation-api/configs"
	"validation-api/pkg/req"
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
	return func(w http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[SendMailRequest](&w, request)
		if err != nil {
			return
		}
		fmt.Print(body)
	}
}

func (handler *verifyHandler) verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
