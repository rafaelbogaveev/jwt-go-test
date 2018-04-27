package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func generateToken(email string, signingKey string) (token Token, err error) {
	now := time.Now().UTC()
	accessToken, err := generateAccessToken([]string{email, now.String()}, signingKey)
	if err != nil {
		return token, err
	}

	// generating refresh token
	refreshToken, err := generateRefreshToken([]string{email, now.String()}, signingKey)
	if err != nil {
		return token, err
	}

	token = Token{
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
		ExpirationDate: now.Add(time.Second * 30),
		Created:        now,
		Updated:        now,
	}

	return token, nil
}


func generateAccessToken(claimArray []string, signingKey string) (encodedToken string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	for i := 0; i < len(claimArray); i++ {
		claims[fmt.Sprintf("claim%d", i+1)] = claimArray[i]
	}

	token.Claims = claims
	encodedToken, err = token.SignedString([]byte(signingKey))

	return encodedToken, err
}

func generateRefreshToken(claimArray []string, signingKey string) (encodedToken string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	for i := 0; i < len(claimArray); i++ {
		claims[fmt.Sprintf("claim%d", i+1)] = "referesh" + claimArray[i]
	}

	token.Claims = claims
	encodedToken, err = token.SignedString([]byte(signingKey))

	return encodedToken, err
}

