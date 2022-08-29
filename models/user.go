package models

import "time"

type User struct{
	Id 	  		string    `json:"id"`
	Email 		string    `json:"email"`
	Pass  		string    `json:"pass"`
	CreatedAt 	time.Time `json:"created_at"`
	Name 		string 	  `json:"name"`
	LastName 	string 	  `json:"last_name"`
	DateBritday time.Time `json:"date_brithday"`
}