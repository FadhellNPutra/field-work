package dto

type ProductsDTO struct {
	ID          string `json:"id"`
	ProductName string `json:"productName,omitempty"`
	Quantity    int `json:"quantity,omitempty"`
	Price       int    `json:"price,omitempty"`
	Material    string `json:"material,omitempty"`
	Description string `json:"description,omitempty"`
}