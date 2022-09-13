package middleware

import (
	"crud-t/means"
	"crud-t/server"
	"net/http"
	"strings"
)



var (
	NO_AUTH_NEEDED = []string{
		"login",
		"singup",
	}
)

func shouldCheckToken(r string) bool  {
	for _, p := range NO_AUTH_NEEDED{
		if strings.Contains(r, p){
			return false
		}
	}
	return true
}


func CheckAuthMiddleware(s server.Server) func (http.Handler) http.Handler  {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			
			if !shouldCheckToken(r.URL.Path){
				next.ServeHTTP(w, r)
				return
			}

			_, err := means.Token(s, w, r)
			if err != nil{
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}