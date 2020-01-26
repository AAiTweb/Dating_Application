package entity

import "github.com/dgrijalva/jwt-go"

var JwtKey = []byte("my_secret_key")

type Claims struct {
	Username       string `json:"username"`
	Id             int    `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	jwt.StandardClaims
}
