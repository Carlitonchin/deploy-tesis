package service

import (
	"crypto/rsa"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/dgrijalva/jwt-go"
)

//-------------------- ID TOKEN ------------------------------

type idTokenClaims struct {
	User *model.User `json:"user"`
	jwt.StandardClaims
}

func generateIDToken(user *model.User, key *rsa.PrivateKey, expiredIn int64) (string, error) {
	now := time.Now().Unix()
	expired_time := now + expiredIn

	claims := &idTokenClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now,
			ExpiresAt: expired_time,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	ss, err := token.SignedString(key)

	return ss, err
}

//-------------------END IDTOKEN ----------------------

// --------------- Refresh token --------------------------
type refreshToken struct {
	ID        string
	SS        string
	ExipresIn time.Duration
}

type refreshTokenClaims struct {
	ID uint
	jwt.StandardClaims
}

func generateRefreshToken(user_id uint, key string, expiresIn int64) (*refreshToken, error) {
	currentTime := time.Now()
	tokenExp := currentTime.Add(time.Duration(expiresIn) * time.Second)

	tokenId := uint(rand.Uint64())

	claims := &refreshTokenClaims{
		ID: user_id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: tokenExp.Unix(),
			Id:        strconv.FormatUint(uint64(tokenId), 10),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(key))

	if err != nil {
		return nil, err
	}

	return &refreshToken{
		ID:        strconv.FormatUint(uint64(tokenId), 10),
		SS:        ss,
		ExipresIn: tokenExp.Sub(currentTime),
	}, nil
}

func validateIdToken(tokenString string, key *rsa.PublicKey) (*idTokenClaims, error) {
	claims := &idTokenClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Token invalido")
	}

	claims, ok := token.Claims.(*idTokenClaims)

	if !ok {
		return nil, fmt.Errorf("No se pudo parsear los claims")
	}

	return claims, nil
}

func validateRefreshToken(refresh_token string, key string) (*refreshTokenClaims, error) {
	claims := &refreshTokenClaims{}

	token, err := jwt.ParseWithClaims(refresh_token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Token inválido")
	}

	claims, ok := token.Claims.(*refreshTokenClaims)

	if !ok {
		return nil, fmt.Errorf("No se pudo parsear las claims del token")
	}

	return claims, nil
}
