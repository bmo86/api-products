package handlers

import (
	"crud-t/means"
	"crud-t/models"
	"crud-t/repository"
	"crud-t/server"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)


const(
	HASH_CONST = 8
)

type SingUpLoginRequest struct{
	Email 		 string `json:"email"`
	Pass  		 string `json:"pass"`
	Name 		 string `json:"name"`
	LastName 	 string `json:"last_name"`
	DateBrithday time.Time `json:"date_brithday"`
}

type SinUpResponse struct{
	Id	  string `json:"id"`
	Email string `json:"email"`
}

type LoginResposne struct{
	Token string `json:"token"`
}

func SingUpHandler(s server.Server) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SingUpLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		hashPasword, err := bcrypt.GenerateFromPassword([]byte(request.Pass), HASH_CONST)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var user =  models.User{
			Email: request.Email,
			Name: request.Name,
			LastName: request.LastName,
			DateBrithday: request.DateBrithday,
			Pass: string(hashPasword),
			Id: id.String(),
		}

		err = repository.NewUser(r.Context(), user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SinUpResponse{
			Id: user.Id,
			Email: user.Email,
		})

	}
}


func HandlerLogin(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req = SingUpLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return 
		}

		user, err := repository.GetByEmailUser(r.Context(), req.Email)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//usuario vacio
		if user == nil {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		}
		//pass no igual
		if err := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(req.Pass)); err != nil{
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}

		claims := models.AppClaims{
			UserId: user.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(2 * time.Hour * 24).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tockenString, err := token.SignedString([]byte(s.Config().JWTSecret))

		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(LoginResposne{
			Token: tockenString,
		})

	}
}


func HandlerMe(s server.Server) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {

		token, err := means.Token(s, w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			user, err := repository.GetByIdUser(r.Context(), claims.UserId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return 
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}


