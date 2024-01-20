package config

const (
	// Group
	ApiGroup      = "/api/v1"
	AdminGroup    = "/admin"
	CustomerGroup = "/customers"

	// Users
	Users     = "users"
	UsersByID = "users/:id"

	// Register
	AdminRegister    = "auth/register/admin"
	CustomerRegister = "auth/register"

	// Login
	CustomerLogin = "auth/login"
	AdminLogin    = "auth/login/admin"
	
	// Products
	Products = "products"
)
