package entity

import "time"

type Transaction struct {
	ID              string    `json:"id"`
	UserId          string    `json:"userId"`
	ProductId       string    `json:"productId"`
	TotalPrice      int       `json:"totalPrice"`
	TotalQuantity   int       `json:"totalQuantity"`
	Size            string    `json:"size"`
	Color           string    `json:"color"`
	Status          string    `json:"status"`
	CustomerMessage string    `json:"customerMessage"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
