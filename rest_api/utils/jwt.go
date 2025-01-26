package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const jwtKey = "my_secret_key_my_secret_key_my_secret_key"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":  email,                                //email
			"userId": userId,                               //user id
			"exp":    time.Now().Add(time.Hour * 2).Unix(), //lifetime of the token
		})

	return token.SignedString([]byte(jwtKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(jwtKey), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("could not parse claims")
	}

	exp := claims["exp"].(float64)
	if time.Now().Unix() > int64(exp) {
		return 0, errors.New("token expired")
	}

	userId := int64(claims["userId"].(float64))

	return userId, nil
}
