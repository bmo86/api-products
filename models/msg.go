package models

type WebSocketMsg struct{
	Type string 		`json:"type"`
	Payload interface{} `json:"payload"`
}