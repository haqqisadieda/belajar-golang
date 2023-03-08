package models

import (
	"encoding/json"
)

type BookCreateRequest struct {
	Title       string      `json:"title,omitempty"       binding:"required"`
	Price       json.Number `json:"price,omitempty"       binding:"required,number"`
	Description string      `json:"description,omitempty" binding:"required"`
	Rating      json.Number `json:"rating,omitempty"      binding:"required,number"`
	Discount    json.Number `json:"discount,omitempty"    binding:"required,number"`
}

type BookUpdateRequest struct {
	Title       string      `json:"title,omitempty"       binding:"required"`
	Price       json.Number `json:"price,omitempty"       binding:"required,number"`
	Description string      `json:"description,omitempty" binding:"required"`
	Rating      json.Number `json:"rating,omitempty"      binding:"required,number"`
	Discount    json.Number `json:"discount,omitempty"    binding:"required,number"`
}
