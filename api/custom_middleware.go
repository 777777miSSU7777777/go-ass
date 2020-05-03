package api

import (
	"context"
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
		if r.Header.Get("authorization") == "" {
			writeError(w, 401, TokenInvalidError, fmt.Errorf("there is no token"))
			return
		}
		tokenParts := strings.Split(r.Header.Get("authorization"), " ")
		var accessToken string
		if len(tokenParts) >= 2 {
			accessToken = tokenParts[1]
		} else {
			writeError(w, 401, TokenInvalidError, fmt.Errorf("token is invalid error"))
			return
		}

		jwtToken, err := jwt.ParseWithClaims(accessToken, &repository.JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(repository.SecretKey), nil
		})

		if jwtToken.Valid {
			payload := jwtToken.Claims.(*repository.JWTPayload)
			ctx := context.WithValue(r.Context(), "userID", payload.ID)
			r = r.WithContext(ctx)
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
