package helper

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	jwt "github.com/dgrijalva/jwt-go"
)

const UserRole = "user"
const AdminRole = "admin"
const AccessTokenDuration = 1800
const RefreshTokenDuration = 518400
const AccessTokenType = "access"
const RefreshTokenType = "refresh"
const SecretKey = "NOT A SECRET KEY"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateTokens(userID int64) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userID,
		"role": UserRole,
		"nbf": time.Now().Add(time.Second * AccessTokenDuration),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"userID": userID,
		"role": UserRole,
		"nbf": time.Now().Add(time.Second * RefreshTokenDuration),
	})
	
	accessTokenString, err := accessToken.SignedString(SecretKey)
	if err != nil {
		return "", "", err
	}

	refreshTokenString, err := refreshToken.SignedString(SecretKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SecretKey), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
