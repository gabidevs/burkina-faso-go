package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type EndPoint struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	Auth     bool
}

func Config(r *mux.Router) *mux.Router {
	points := EndPoints

	for _, point := range points {
		r.HandleFunc(point.URI, point.Function).Methods(point.Method)
	}

	return r

}

func Generate() *mux.Router {
	r := mux.NewRouter()

	return Config(r)

}
