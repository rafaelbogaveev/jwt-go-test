package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func addPrivateRoutes(r *mux.Router) {
	r.HandleFunc("/test", authorizedOnly(Test)).Methods("GET")
}

func Test(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tokenString := r.FormValue("x-auth")

	w.Write([]byte(tokenString))
}

func authorizedOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.FormValue("email")
		userToken := r.Header.Get("x-auth")

		token, ok := tokens[email];
		if !ok {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		if token.ExpirationDate.Before(time.Now().UTC()) {
			http.Error(w, http.StatusText(403), 403)
			return
		}

		encodedToken, err := generateAccessToken([]string{email, token.Updated.String()}, signingKey)
		if err != nil {
			println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		if encodedToken != userToken {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		next(w, r)
	}
}
