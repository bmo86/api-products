package models


type Product struct{
	
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	StockMin    int     `json:"stockMin"`
	Description string  `json:"description"`
	UserId 		string  `json:"user_id"`
}