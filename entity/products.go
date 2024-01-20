package entity

import (
  _ "field_work/helpers"
  "time"
)

type Products struct {
	ID          string `json:"id"`
	ProductName string `json:"productName,omitempty"`
	Quantity    int `json:"quantity,omitempty"`
	Price       int    `json:"price,omitempty"`
	Material    string `json:"material,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

func (p *Products) TimeFormat(fields ...string) {
  for _, field := range fields {
    if field != "" {
      switch field {
      case "CreatedAt":
        createdAt, _ := time.Parse("2006-01-02T15:04:05+07:00", p.CreatedAt)
        p.CreatedAt = createdAt.Format(time.RFC850)
      case "UpdatedAt":
        updatedAt, _ := time.Parse("2006-01-02T15:04:05+07:00", p.UpdatedAt)
        p.UpdatedAt = updatedAt.Format(time.RFC850)
      }
    }
  }
}
