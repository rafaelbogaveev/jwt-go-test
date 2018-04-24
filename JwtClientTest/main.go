package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main()  {
	r := mux.NewRouter()
	addRoutes(r)

	http.ListenAndServe(":3003", r)
}