package models

type History struct {
	ID        int    `json:"id"`
	Customer  string `json:"customer"`
	Action    string `json:"action"`
	Timestamp string `json:"timestamp"`
}
