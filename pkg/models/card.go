package models

type Card struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	ListID string `json:"idList"`
	Status string `json:"-"`
}
