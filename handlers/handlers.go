package handlers

import (
	"net/http"
)

func HomeHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("access-control-allow-origin", "*")
		writer.Header().Set("access-control-allow-methods", "GET, POST, PATCH, DELETE")
		writer.Header().Set("access-control-allow-headers", "accept, content-type")

		if (*request).Method == "OPTIONS" {
			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}
