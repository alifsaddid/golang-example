package models

type Token struct {
	Token        string
	TokenJson    string
	UserId       int
	ClientId     string
	ClientSecret string
	IssuedAt     int64
	Type         string
	RefreshToken string
}
