package verify

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"validation-api/configs"
	"validation-api/pkg/req"
	"validation-api/pkg/res"

	"github.com/jordan-wright/email"
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

		hash := fmt.Sprintf("%x", rand.Uint64())

		data := map[string]string{
			"email": body.Email,
			"hash":  hash,
		}

		jsonData, _ := json.Marshal(data)
		os.WriteFile(filepath.Join("verify.json"), jsonData, 0644)

		e := email.NewEmail()
		e.From = handler.conf.Email
		e.To = []string{body.Email}
		e.Subject = "Email Verification"
		e.Text = []byte(fmt.Sprintf("Для подтверждения перейди по ссылке:\nhttp://localhost:8081/verify/%s", hash))
		err = e.Send(handler.conf.Address+":587", smtp.PlainAuth("", handler.conf.Email, handler.conf.Password, handler.conf.Address))

		if err != nil {
			res.Json(w, ErrorResponse{
				Message: err.Error(),
			}, 500)
			return
		}

		res.Json(w, "", 200)
	}
}

func (handler *verifyHandler) verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")
		data := map[string]string{}
		content, err := os.ReadFile("verify.json")

		if err != nil {
			log.Printf("Ошибка чтения verify.json: %v", err)
			res.Json(w, ErrorResponse{
				Message: "Не удалось найти код верификации",
			}, 404)
			return
		}
		json.Unmarshal(content, &data)

		if data["hash"] == hash {
			res.Json(w, "", 200)
		} else {
			log.Printf("Проверка hash: %s, ожидается: %s", hash, data["hash"])
			res.Json(w, ErrorResponse{
				Message: "Некорректный код подтверждения",
			}, 400)
		}

		os.Remove("verify.json")
	}
}
