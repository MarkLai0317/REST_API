package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct {
	router *chi.Mux
}

// var (
// 	chiDispathcher = chi.NewRouter()
// )

func NewChiRouter() *chiRouter {
	return &chiRouter{
		router: chi.NewRouter(),
	}
}

func (router *chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.router.Get(uri, f)
}
func (router *chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.router.Post(uri, f)
}

func (router *chiRouter) SERVE(port string) {
	fmt.Printf("Chi server run on port: %v", port)
	http.ListenAndServe(port, router.router)
}
