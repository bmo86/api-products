package models


type Product struct{
	
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	StockMin    int     `json:"stockMin"`
	Description string  `json:"description"`

}