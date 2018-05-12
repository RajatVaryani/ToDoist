package router

import (
	"github.com/gorilla/mux"
	"ToDoist/handlers"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomeHandler()).Methods(http.MethodGet)
	return router
}
