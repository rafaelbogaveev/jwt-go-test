package main

import (
	"testing"
	"net/http"
	"io/ioutil"
	"fmt"
	"time"
	"encoding/json"
)

func TestRequest(t *testing.T) {
	email := "test@test.ru"
	accessToken := ""
	refreshToken := ""

	for {
		time.Sleep(1000 * time.Millisecond)

		start := time.Now()

		if len(accessToken) == 0 {
			token, err := getNewToken(email)
			if err != nil{
				println(err)
				continue
			}

			accessToken = token.AccessToken
			refreshToken = token.RefreshToken
			println(fmt.Sprintf("new access token %s, new refresh token %s", accessToken, refreshToken))
		}

		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:3002/test?email=%s", email), nil)

		if err != nil {
			println(err)
			fmt.Println(time.Since(start))
			continue
		}

		req.Header.Set("x-auth", accessToken)
		resp, err := client.Do(req)

		if err != nil {
			println(err)
			fmt.Println(time.Since(start))
			continue
		}

		// exception occured during request
		if resp.StatusCode == 500 {
			println("Internal Server Error")
			fmt.Println(time.Since(start))
			continue
		}

		// request unauthorized, token is empty or incorrect
		if resp.StatusCode == 401 {
			println("Unathorized")
			fmt.Println(time.Since(start))
			continue
		}

		// it is forbidden to get access by current token
		if resp.StatusCode == 403 {
			println("token expired")
			token, err := getRefreshToken(email)
			if err != nil{
				accessToken = ""
				println(err)
				fmt.Println(time.Since(start))
				continue
			}

			accessToken = token.AccessToken
			refreshToken = token.RefreshToken
			println(fmt.Sprintf("new access token %s, new refresh token %s", accessToken, refreshToken))
			fmt.Println(time.Since(start))
			continue
		}

		// request was successfully processed
		if resp.StatusCode == 200 {
			println("ok")
			fmt.Println(time.Since(start))
			continue
		}
	}

}

func getNewToken(email string) (token Token, err error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:3002/register?email=%s", email))

	if err != nil {
		return token, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return token, err
	}

	err = json.Unmarshal(data, &token);
	return token, err
}

func getRefreshToken(email string) (token Token, err error)  {
	resp, err := http.Get(fmt.Sprintf("http://localhost:3002/refresh?email=%s", email))

	if err != nil {
		return token, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return token, err
	}

	err = json.Unmarshal(data, &token);
	return token, err
}