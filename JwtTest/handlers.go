package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
)

func addRoutes(r *mux.Router) {
	r.HandleFunc("/test", authorizedOnly(Test)).Methods("GET")
}

func Test(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tokenString := r.FormValue("token")

	w.Write([]byte(tokenString))
}

func authorizedOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//email := r.FormValue("email")
		tokenString := r.FormValue("token")

		token := jwt.New(jwt.SigningMethodHS256)
		encodedToken, err := token.SignedString([]byte(signingKey))

		if err != nil {
			println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		if tokenString != encodedToken {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		next(w, r)
	}
}
