package model

type OAuthInfo struct {
	Username string `json:"login"`
	OAuthId  int64  `json:"id"`
	Avatar   string `json:"avatar_url"`
}
