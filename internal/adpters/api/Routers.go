package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

var router *chi.Mux

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
}

func Routers() *chi.Mux {
	router.Get("/api/trans", GetFunction)
	router.Post("/api/trans", Createtransaction)

	return router
}
func Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r) // dispatch the request
	})
}
