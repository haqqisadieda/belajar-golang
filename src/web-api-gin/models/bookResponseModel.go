package models

type BookResponse struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Price       int    `json:"price,omitempty"`
	Rating      int    `json:"rating,omitempty"`
	Discount    int    `json:"discount,omitempty"`
}
