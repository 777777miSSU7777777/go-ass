package helper

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)

type UploadTrackCallback func(string) error

const (
	AccessTokenType = "access"
	RefreshTokenType = "refresh"
	DefaultSecretKey = "NOT A SECRET KEY"
	DefaultAccessExp = 1200
	DefaultRefreshExp = 5184000
)

func SignToken(userID string, tokenType string, exp int, secretKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["token_type"] = tokenType
	claims["exp"] = time.Now().Add(time.Second * time.Duration(exp)).Unix()

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("sign token error: %v", err)
	}

	return signedToken, nil
}

func VerifyToken(tokenString string, secretKey string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}