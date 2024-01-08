package config

const (
	// Users query
	InsertUser                   = "INSERT INTO users (name, username, password, address, role) VALUES ($1, $2, crypt($3, gen_salt('md5')), $4, $5) RETURNING id, name, username, address, role, created_at;"
	SelectAllUser                = "SELECT id, name, username, address, role, created_at FROM users ORDER BY role DESC, created_at DESC LIMIT $1 OFFSET $2;"
	SelectUserByID               = "SELECT id, name, username, address, role FROM users WHERE id = $1;"
	SelectUserByUsername         = "SELECT id, name, username, address, role FROM users WHERE username ILIKE $1;"
	SelectUserByUsernameForLogin = "SELECT id, name, username, password, role FROM users WHERE username = $1 AND password = crypt($2, password);"
	UpdateUser                   = "UPDATE users SET name = $1, username = $2, password = crypt($3, password), address = $4, role = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $6 RETURNING created_at, updated_at;"
	DeleteUser                   = "DELETE FROM users WHERE id = $1;"
)
