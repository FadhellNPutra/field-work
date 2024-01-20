package entity

import (
  "fmt"
  _ "field_work/helpers"
  "reflect"
  "strconv"
  "time"
)

type Products struct {
	ID          string `json:"id"`
	ProductName string `json:"productName,omitempty"`
	Quantity    int64    `json:"quantity,omitempty"`
	Price       int64    `json:"price,omitempty"`
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

func (p *Products) ToMap() map[string]*string {
	structValue := reflect.ValueOf(*p)
	structType := structValue.Type()

	result := make(map[string]*string)

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldName := structType.Field(i).Name
		
		switch field.Kind() {
		case reflect.String:
  		fieldValue := field.String()
  		result[fieldName] = &fieldValue
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
  		fieldValue := field.Int()
  		fieldValueStr := strconv.FormatInt(fieldValue, 10)
  		result[fieldName] = &fieldValueStr
		default:
			fmt.Printf("unhandled kind %s\n", structValue.Kind())
		}
	}

	return result
}