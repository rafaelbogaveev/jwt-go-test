package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Token struct {
	ID             bson.ObjectId `json:"id" bson:"_id,omitempty"`
	AccessToken    string        `json:"accessToken" bson:"accessToken"`
	RefreshToken   string        `json:"refreshToken" bson:"refreshToken"`
	ExpirationDate time.Time     `json:"expirationDate" bson:"expirationDate"`
	Created        time.Time     `json:"created" bson:"created"`
	Updated        time.Time     `json:"updated" bson:"updated"`
}
