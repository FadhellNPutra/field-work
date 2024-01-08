package entity

type Transaction struct {
	ID              string `json:"id"`
	UserId          string `json:"userId,omitempty"`
	ProductId       string `json:"productId,omitempty"`
	TotalPrice      int    `json:"totalPrice,omitempty"`
	TotalQuantity   int    `json:"totalQuantity,omitempty"`
	Size            string `json:"size,omitempty"`
	Color           string `json:"color,omitempty"`
	Status          string `json:"status,omitempty"`
	CustomerMessage string `json:"customerMessage,omitempty"`
	CreatedAt       string `json:"createdAt,omitempty"`
	UpdatedAt       string `json:"updatedAt,omitempty"`
}
