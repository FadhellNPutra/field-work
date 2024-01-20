package entity

import (
  "strings"
)

type ProductImages struct {
  ID        string `json:"id"`
  ProductID string `json:"productId,omitempty"`
  FileName  string `json:"fileName,omitempty"`
  IsPrimary bool   `json:"isPrimary,omitempty"`
}

func (p *ProductImages) IsAllowedExtension(ext string, allowedExtensions ...string) bool {
  for _, allowedExtension := range allowedExtensions {
		if strings.ToLower(ext) == allowedExtension {
			return true
		}
	}

	return false
}