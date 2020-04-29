package api

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/777777miSSU7777777/go-ass/repository"
)

var TokenInvalidError = "TOKEN INVALID ERROR"
var TokenExpiredError = "TOKEN EXPIRED ERROR"

func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := strings.Split(r.Header.Get("authorization"), " ")
		var accessToken string
		if len(bearerToken) >= 2 {
			accessToken = bearerToken[1]
		} else {
			writeError(w, 401, TokenInvalidError, fmt.Errorf("token is invalid error"))
			return
		}

		jwtToken, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(repository.SecretKey), nil
		})

		if jwtToken.Valid {
			next.ServeHTTP(w, r)
		} else if validationError, ok := err.(*jwt.ValidationError); ok {
			if validationError.Errors&jwt.ValidationErrorMalformed != 0 {
				writeError(w, 401, TokenInvalidError, fmt.Errorf("token is invalid error"))
				return
			} else if validationError.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				writeError(w, 401, TokenExpiredError, fmt.Errorf("token is expired"))
				return
			} else {
				writeError(w, 401, TokenInvalidError, fmt.Errorf("token is invalid error"))
				return
			}
		} else {
			writeError(w, 401, TokenInvalidError, fmt.Errorf("token is invalid error"))
			return
		}
	})
}