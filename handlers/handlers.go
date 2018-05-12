package handlers

import "net/http"

func HomeHandler() http.HandlerFunc  {
	return func(writer http.ResponseWriter, request *http.Request) {
		enableCors(&writer)
		writer.WriteHeader(http.StatusOK)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
