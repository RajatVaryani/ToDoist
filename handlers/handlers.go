package handlers

import "net/http"

func HomeHandler() http.HandlerFunc  {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}
}
