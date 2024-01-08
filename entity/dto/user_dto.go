package dto

type UserDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}

type UpdateUserDTO struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}
