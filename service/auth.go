package service

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type key int

const (
	ContextAuthUIDKey key = iota
	ContextAuthIsAuthedKey
)

type AuthService struct {
	PubKey  string
	PrivKey string
	Expire  time.Duration // in hours
}

type JwtClaims struct {
	UserID int
	*jwt.StandardClaims
}

func (a *AuthService) getPublicKey() (*rsa.PublicKey, error) {
	keyData, err := ioutil.ReadFile(a.PubKey)

	if err != nil {
		return nil, err
	}

	parsedKey, err := jwt.ParseRSAPublicKeyFromPEM(keyData)

	return parsedKey, err
}

func (a *AuthService) getPrivateKey() (*rsa.PrivateKey, error) {
	keyData, err := ioutil.ReadFile(a.PrivKey)

	if err != nil {
		return nil, err
	}

	parsedKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)

	return parsedKey, err
}

// CreateToken jwt
func (a *AuthService) CreateToken(userID int) (string, error) {
	k, err := a.getPrivateKey()
	if err != nil {
		return "", err
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, a.getClaims(userID))

	return t.SignedString(k)
}

func (a *AuthService) CheckToken(ts string) (int, error) {
	t, err := jwt.ParseWithClaims(ts, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		key, err := a.getPublicKey()

		return key, err
	})

	if err != nil {
		return 0, err
	}

	if !t.Valid {
		return 0, errors.New("Token is invalid")
	}

	claims, ok := t.Claims.(*JwtClaims)
	if !ok {
		return 0, errors.New("Invalid claim")
	}

	return claims.UserID, err
}

func (a *AuthService) getClaims(uid int) *JwtClaims {
	expireTime := time.Now().Add(time.Hour * a.Expire).Unix()
	claims := &JwtClaims{
		UserID:         uid,
		StandardClaims: &jwt.StandardClaims{ExpiresAt: expireTime},
	}

	return claims
}

func (a *AuthService) Middleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		uid, err := a.CheckToken(token)

		isAuthed := false
		if err == nil {
			isAuthed = true
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, ContextAuthIsAuthedKey, isAuthed)
		ctx = context.WithValue(ctx, ContextAuthUIDKey, uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
