package entity

type Products struct {
	ID          string `json:"id"`
	ProductName string `json:"productName,omitempty"`
	Quantity    string `json:"quantity,omitempty"`
	Price       int    `json:"price,omitempty"`
	Material    string `json:"material,omitempty"`
	Description string `json:"description,omitempty"`
	Photo       string `json:"photo,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}
