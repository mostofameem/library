package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"user_service/config"
	"user_service/db"
	"user_service/web/utils"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
	jwt.RegisteredClaims
}

// Define a custom type for the context key
func GenerateToken(usr db.User) (string, error) {
	conf := config.GetConfig()
	expirationTime := time.Now().Add(60 * time.Minute).Unix()

	accessToken, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"Id":   usr.Id,
			"Type": usr.Type,
			"exp":  expirationTime,
		},
	).SignedString([]byte(conf.JwtSecret))
	if err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("error")
	}
	return accessToken, nil
}

func unauthorizedResponse(w http.ResponseWriter) {
	utils.SendError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
}
func VerifyToken(tokenStr string) (AuthClaims, error) {
	conf := config.GetConfig()
	var claims = AuthClaims{}
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(conf.JwtSecret), nil
		},
	)

	if !token.Valid {
		err = fmt.Errorf("unauthorized")
	}
	return claims, err
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// collect token from header
		header := r.Header.Get("authorization")
		tokenStr := ""

		// collect token from query
		if len(header) == 0 {
			tokenStr = r.URL.Query().Get("auth")
		} else {
			tokens := strings.Split(header, " ")
			if len(tokens) != 2 {
				unauthorizedResponse(w)
				return
			}
			tokenStr = tokens[1]
		}
		claims, err := VerifyToken(tokenStr)

		// set user id in the context
		if err != nil {
			unauthorizedResponse(w)
			return
		}
		r.Header.Set("id", strconv.Itoa(claims.Id))
		next.ServeHTTP(w, r)
	})
}

func GetUserIDFromToken(tokenStr string) (int, error) {
	conf := config.GetConfig()

	// Parse JWT
	var claims AuthClaims
	_, err := jwt.ParseWithClaims(
		tokenStr,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(conf.JwtSecret), nil
		},
	)
	if err != nil {
		return 0, err
	}

	// Return user ID from claims
	return claims.Id, nil
}
func GetUserUserTypeFromToken(tokenStr string) (string, error) {
	conf := config.GetConfig()

	// Parse JWT
	var claims AuthClaims
	_, err := jwt.ParseWithClaims(
		tokenStr,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(conf.JwtSecret), nil
		},
	)
	if err != nil {
		return "", err
	}

	// Return user ID from claims
	return claims.Type, nil
}
