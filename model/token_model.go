package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	User User
	Type string // REFRESH_TOKEN and TOKEN 用于更新token的标识
	Time time.Time
	jwt.StandardClaims
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}
