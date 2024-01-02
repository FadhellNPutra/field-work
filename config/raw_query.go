package config

const(
	// Users query
	InsertUser = "INSERT INTO users (name, username, password, address, role, updated_at) VALUES ($1, $2, crypt($3, gen_salt('md5')), $4, $5, CURRENT_TIMESTAMP) RETURNING id, created_at, updated_at;"
	SelectAllUser = "SELECT id, name, username, password, address, role, created_at, updated_at FROM users LIMIT $1 OFFSET $2;"
	SelectUserByID = "SELECT id, name, username, password, address, role, created_at, updated_at FROM users WHERE id = $1;"
	SelectUserByUsername = "SELECT id, name, username, password, address, role, created_at, updated_at FROM users WHERE username = $1;"
	SelectUserByUsernameForLogin = "SELECT id, name, username, password, address, role, created_at, updated_at FROM users WHERE username = $1, password = $2;"
	UpdateUser = "UPDATE users SET name = $1, username = $2, password = crypt($3, password), address = $4, role = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $8 RETURNING created_at, updated_at"
	DeleteUser = "DELETE FROM users WHERE id = $1;"
)