package model

type Meta struct {
	Status    string `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type SingleResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type PagedResponse struct {
	Meta   Meta          `json:"meta"`
	Data   []interface{} `json:"data"`
	Paging Paging        `json:"paging,omitempty"`
}
