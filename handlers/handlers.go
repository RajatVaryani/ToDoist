package handlers

import "net/http"

func HomeHandler() http.HandlerFunc  {
	return func(writer http.ResponseWriter, request *http.Request) {
		setupResponse(&writer, request)

		if (*request).Method == "OPTIONS" {
			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("access-control-allow-origin", "*")
	(*w).Header().Set("access-control-allow-methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("access-control-allow-headers", "accept, content-Type")
}
