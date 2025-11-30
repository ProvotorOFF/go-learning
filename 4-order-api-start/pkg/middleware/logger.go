package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type ResponseWriterWrapper struct {
	http.ResponseWriter
	Status int
}

func (w *ResponseWriterWrapper) WriteHeader(status int) {
	w.Status = status
	w.ResponseWriter.WriteHeader(status)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := &ResponseWriterWrapper{
			w,
			200,
		}

		next.ServeHTTP(ww, r)

		log := logrus.New()
		log.SetFormatter(&logrus.JSONFormatter{})
		log.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
			"status": ww.Status,
		}).Info("request completed")
	})
}
