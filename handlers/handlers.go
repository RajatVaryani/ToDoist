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
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
