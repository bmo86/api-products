package handlers

import (
	"crud-t/server"
	"encoding/json"
	"net/http"
)

type HomeResponse struct{
	Msg string  `json:"msg"`
	Status bool `json:"status"`
}

func HandlerHome(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeResponse{
			Msg: "Welcome to server products",
			Status: true,
		})
	}
}
