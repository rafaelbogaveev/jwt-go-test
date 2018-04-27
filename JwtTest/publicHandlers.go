package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func addPublicRoutes(r *mux.Router) {
	r.HandleFunc("/register", register).Methods("GET")
	r.HandleFunc("/refresh", refreshToken).Methods("GET")
}

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")

	if len(email) == 0 {
		println("email is empty")
		http.Error(w, http.StatusText(500), 500)
		return
	}

	token, err := generateToken(email, signingKey)
	if err != nil {
		println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// save token in memory
	tokens[email] = token

	// sending via JSON
	response, err := json.Marshal(token)
	if err != nil {
		println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	println(fmt.Sprintf("New access token=%s \n New refresh token=%s \n ExpirationDate=%s", token.AccessToken, token.RefreshToken, token.ExpirationDate.String()))

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}


func refreshToken(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")

	oldToken, ok := tokens[email]

	if !ok {
		println("No token to refresh")
		http.Error(w, http.StatusText(500), 500)
		return
	}

	token, err := generateToken(email, signingKey)
	if err != nil {
		println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	token.Created = oldToken.Created
	tokens[email] = token

	// sending via JSON
	response, err := json.Marshal(token)
	if err != nil {
		println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	println(fmt.Sprintf("Refreshed access token=%s\n New refresh token=%s\n ExpirationDate=%s", token.AccessToken, token.RefreshToken, token.ExpirationDate.String()))

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
