package entity

import "time"

type Products struct {
	ID          string    `json:"id"`
	ProductName string    `json:"productName"`
	Quantity    string    `json:"quantity"`
	Price       int       `json:"price"`
	Material    string    `json:"material"`
	Description string    `json:"description"`
	Photo       string    `json:"photo"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}