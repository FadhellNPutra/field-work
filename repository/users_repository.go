package repository

import (
	"database/sql"
	"field_work/config"
	"field_work/entity"
	"field_work/shared/model"
	"log"
	"math"
	"time"
)

type UsersRepository interface {
	List(page, size int) ([]entity.Users, model.Paging, error)
	GetUsersByID(id string) (entity.Users, error)
	GetUsersByUsername(username string) (entity.Users, error)
	GetUsersByUsernameForLogin(username, password string) (entity.Users, error)
	CreateUser(payload entity.Users) (entity.Users, error)
	UpdateUser(payload entity.Users) (entity.Users, error)
	DeleteUser(id string) error
}

type usersRepository struct {
	db *sql.DB
}

func (u *usersRepository) CreateUser(payload entity.Users) (entity.Users, error) {
	var users entity.Users

	err := u.db.QueryRow(config.InsertUser,
		payload.Name,
		payload.Username,
		payload.Password,
		payload.Address,
		payload.Role,
	).
		Scan(
			&users.ID,
			&users.Name,
			&users.Username,
			&users.Address,
			&users.Role,
			&users.CreatedAt,
		)

	if err != nil {
		log.Println("usersRepository.QueryRow: ", err.Error())
		return entity.Users{}, err
	}

	return users, nil
}

func (u *usersRepository) DeleteUser(id string) error {
	err := u.db.QueryRow(config.DeleteUser, id).Err()
	if err != nil {
		log.Println("usersRepository.DeleteUser.QueryRow: ", err.Error())
		return err
	}

	log.Println(err)
	return nil
}

func (u *usersRepository) GetUsersByID(id string) (entity.Users, error) {
	var users entity.Users
	err := u.db.QueryRow(config.SelectUserByID, id).Scan(
		&users.ID,
		&users.Name,
		&users.Username,
		&users.Address,
		&users.Role,
	)
	if err != nil {
		log.Println("usersRepository.GetUsersByID.QueryRow: ", err.Error())
		return entity.Users{}, err
	}
	return users, nil
}

func (u *usersRepository) GetUsersByUsername(username string) (entity.Users, error) {
	var users entity.Users
	err := u.db.QueryRow(config.SelectUserByUsername, username).Scan(
		&users.ID,
		&users.Name,
		&users.Username,
		&users.Address,
		&users.Role,
	)
	if err != nil {
		log.Println("usersRepository.GetUsersByUsername.QueryRow: ", err.Error())
		return entity.Users{}, err
	}
	return users, nil
}

func (u *usersRepository) GetUsersByUsernameForLogin(username string, password string) (entity.Users, error) {
	var users entity.Users
	err := u.db.QueryRow(config.SelectUserByUsernameForLogin, username, password).Scan(
		&users.ID,
		&users.Name,
		&users.Username,
		&users.Password,
		&users.Role,
	)
	if err != nil {
		log.Println("usersRepository.GetUsersByUsername.QueryRow: ", err.Error())
		return entity.Users{}, err
	}
	return users, nil
}

func (u *usersRepository) List(page int, size int) ([]entity.Users, model.Paging, error) {
	var users []entity.Users
	offset := (page - 1) * size
	rows, err := u.db.Query(config.SelectAllUser, size, offset)
	if err != nil {
		log.Println("usersRepository.Query: ", err.Error())
		return nil, model.Paging{}, err
	}
	for rows.Next() {
		var user entity.Users
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Address,
			&user.Role,
			&user.CreatedAt,
		)
		if err != nil {
			log.Println("usersRepository.rows.Next(): ", err.Error())
			return nil, model.Paging{}, err
		}
		users = append(users, user)
	}

	totalRows := 0
	if err := u.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&totalRows); err != nil {
		return nil, model.Paging{}, err
	}

	paging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}
	return users, paging, nil
}

func (u *usersRepository) UpdateUser(payload entity.Users) (entity.Users, error) {
	var users entity.Users
	payload.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	err := u.db.QueryRow(config.UpdateUser,
		payload.Name,
		payload.Username,
		payload.Password,
		payload.Address,
		payload.Role,
		payload.ID,
	).Scan(&users.CreatedAt, &users.UpdatedAt)
	if err != nil {
		log.Println("usersRepository.QuerRow: ", err.Error())
		return entity.Users{}, err
	}
	users.ID = payload.ID
	users.Name = payload.Name
	users.Username = payload.Username
	users.Address = payload.Address
	users.Role = payload.Role

	return users, nil

}

func NewUsersRepository(db *sql.DB) UsersRepository {
	return &usersRepository{db: db}
}
