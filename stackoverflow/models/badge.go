package models

type Badge struct {
	ID          unit   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
