package handlers

import (
	"crud-t/means"
	"crud-t/models"
	"crud-t/repository"
	"crud-t/server"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)


type ProductRequest struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	StockMin    int     `json:"stockMin"`
	Description string  `json:"description"`

}

type MsgProduct struct{
	Msg 		string 	`json:"msg"`
	Id 			string 	`json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	StockMin    int     `json:"stockMin"`
	Description string  `json:"description"`

}


func NewProductHandler( s server.Server ) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token, err := means.Token(s, w, r)

		if err != nil{
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			var request = ProductRequest{}
			
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var product = models.Product{
				Name: 		 request.Name,
				Price: 		 request.Price, 
				Stock: 		 request.Stock,
				StockMin: 	 request.StockMin,
				Description: request.Description,
				Id: id.String(),
				UserId: claims.UserId,
			}

			err = repository.NewProduct(r.Context(), &product)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var productMsgWs = models.WebSocketMsg{
				Type: "Product Creted!",
				Payload: product,
			}

			s.Hub().BroadCast(productMsgWs, nil)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(MsgProduct{
				Msg: 		"Created Product!",
				Id: 		 product.Id,
				Name: 		 product.Name,
				Price: 		 product.Price, 
				Stock: 		 product.Stock,
				StockMin: 	 product.StockMin,
				Description: product.Description,

			})

		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
	}
}

func UpdateProductHandler( s server.Server ) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		idP := mux.Vars(r)
		
		token, err := means.Token(s, w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			var productResquest = ProductRequest{}

			if err := json.NewDecoder(r.Body).Decode(&productResquest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			updateProduct := models.Product{
				Id: idP["productID"],
				Name: 			productResquest.Name,
				Price: 			productResquest.Price,
				Stock: 			productResquest.Stock,
				StockMin: 		productResquest.StockMin,
				Description: 	productResquest.Description,
				UserId: 		claims.UserId,	
			}

			err = repository.UpdateProduct(r.Context(), &updateProduct)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var MsgProductWs = models.WebSocketMsg{
				Type: "Product Update",
				Payload: updateProduct,
			}
			s.Hub().BroadCast(MsgProductWs, nil)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(MsgProduct{
				Msg: 		"Update Product!",
				Id: 		 idP["producID"],
				Name: 		 productResquest.Name,
				Price: 		 productResquest.Price, 
				Stock: 		 productResquest.Stock,
				StockMin: 	 productResquest.StockMin,
				Description: productResquest.Description,
			})

		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}









