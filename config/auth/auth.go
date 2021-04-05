package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/KenLu0929/flowControlTask/config"
)

type Claims struct {
	Account  string `json:"account"`
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtSecret = []byte(config.Config.JWT.Secret)

func SignToken(account string, id uint, username string) (string, error) {
	now := time.Now()
	jwtId := account + username + strconv.FormatInt(now.Unix(), 10)

	//set claims
	claims := Claims{
		Account:  account,
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Audience:  account,
			ExpiresAt: now.Add(time.Hour * 1).Unix(),
			Id:        jwtId,
			IssuedAt:  now.Unix(),
			NotBefore: now.Add(time.Second * 0).Unix(),
			Subject:   account,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

func AuthRequired(token string) (account string, id uint, username string, err error) {
	var message string

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (secret interface{}, err error) {
		return jwtSecret, nil
	})

	if err != nil {
		message = err.Error()

		if validationErr, ok := err.(*jwt.ValidationError); ok {
			switch validationErr.Errors {
			case jwt.ValidationErrorMalformed:
				message = "Token is malformed"
			case jwt.ValidationErrorUnverifiable:
				message = "Token could not be verified because of signing problems"
			case jwt.ValidationErrorSignatureInvalid:
				message = "Signature validation failed"
			case jwt.ValidationErrorExpired:
				message = "Exp validation failed"
			case jwt.ValidationErrorNotValidYet:
				message = "NBF validation failed"
			default:
				message = "can not handle this token"
			}
		}

		return "", 0, "", errors.New(message)
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok {
		if claims.Account == "" || claims.ID == 0 || claims.Username == "" {
			return "", 0, "", errors.New("JWT token payload is improper")
		} else {
			return claims.Account, claims.ID, claims.Username, nil
		}
	} else {
		return "", 0, "", errors.New("JWT token payload is improper")
	}
}
