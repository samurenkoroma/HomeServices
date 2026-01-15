package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"samurenkoroma/services/configs"
	"time"
)

type JWTData struct {
	Email string
}

type Tokens struct {
	Access  string
	Refresh string
}
type JWT struct {
	AccessSecret  string
	RefreshSecret string
}

func NewJWT(secrets configs.AuthConfig) *JWT {
	return &JWT{
		AccessSecret:  secrets.AccessSecret,
		RefreshSecret: secrets.RefreshSecret,
	}
}

func (j *JWT) Create(data JWTData) (Tokens, error) {
	accessClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Minute).Unix(),
	})
	access, err := accessClaims.SignedString([]byte(j.AccessSecret))
	if err != nil {
		return Tokens{}, err
	}
	refreshClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * 24 * 72).Unix(),
	})
	refresh, err := refreshClaims.SignedString([]byte(j.RefreshSecret))
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{
		Refresh: refresh,
		Access:  access,
	}, nil
}

func (j *JWT) ParseAccess(token string) (bool, *JWTData) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(j.AccessSecret), nil
	})

	if err != nil {
		return false, nil
	}

	email := t.Claims.(jwt.MapClaims)["email"]
	return t.Valid, &JWTData{
		Email: email.(string),
	}
}

func (j *JWT) ParseRefresh(token string) (bool, *JWTData) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(j.RefreshSecret), nil
	})

	if err != nil {
		return false, nil
	}

	email := t.Claims.(jwt.MapClaims)["email"]
	return t.Valid, &JWTData{
		Email: email.(string),
	}
}
