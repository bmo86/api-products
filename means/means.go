package means

import (
	"crud-t/models"
	"crud-t/server"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func Token(s server.Server, w http.ResponseWriter, r *http.Request) (*jwt.Token, error)  {
	tokenstring := strings.TrimSpace(r.Header.Get("Authorization"))

	token, err := jwt.ParseWithClaims(tokenstring, &models.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})

	return token, err
}