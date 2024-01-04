package usecase

import (
	"errors"
	"field_work/entity"
	"field_work/repository"
	"field_work/shared/model"
	"fmt"
)

type UsersUseCase interface {
	FindUsersByID(id string) (entity.Users, error)
	FindUsersByUsername(username string) (entity.Users, error)
	FindUsersForLogin(username, password string) (entity.Users, error)
	RegisterNewUsers(payload entity.Users) (entity.Users, error)
	UpdateUsers(payload entity.Users) (entity.Users, error)
	DeleteUsers(id string) error
	ListAll(page, size int) ([]entity.Users, model.Paging, error)
}

type usersUseCase struct {
	repo repository.UsersRepository
}

// DeleteUsers implements UsersUseCase.
func (u *usersUseCase) DeleteUsers(id string) error {
	if id == "" {
		return errors.New("gagal menghapus data")
	}
	return u.repo.DeleteUser(id)
}

// FindUsersByID implements UsersUseCase.
func (u *usersUseCase) FindUsersByID(id string) (entity.Users, error) {
	if id == ""{
		return entity.Users{}, errors.New("id harus diisi")
	}
	return u.repo.GetUsersByID(id)
}

// FindUsersByUsername implements UsersUseCase.
func (u *usersUseCase) FindUsersByUsername(username string) (entity.Users, error) {
	if username == ""{
		return entity.Users{}, errors.New("username harus diisi")
	}
	return u.repo.GetUsersByUsername(username)
}

// FindUsersForLogin implements UsersUseCase.
func (u *usersUseCase) FindUsersForLogin(username string, password string) (entity.Users, error) {
	if username == "" || password == ""{
		return entity.Users{}, errors.New("gagal masuk")
	}
	return u.repo.GetUsersByUsernameForLogin(username, password)
}

// ListAll implements UsersUseCase.
func (u *usersUseCase) ListAll(page int, size int) ([]entity.Users, model.Paging, error) {
	return u.repo.List(page, size)
}

// RegisterNewUsers implements UsersUseCase.
func (u *usersUseCase) RegisterNewUsers(payload entity.Users) (entity.Users, error) {
	user, err := u.repo.CreateUser(payload)
	if err != nil{
		return entity.Users{}, fmt.Errorf("failed save data User: %v", err.Error())
	}
	return user, nil
}

// UpdateUsers implements UsersUseCase.
func (u *usersUseCase) UpdateUsers(payload entity.Users) (entity.Users, error) {
	user, err := u.repo.UpdateUser(payload)
	if err != nil{
		return entity.Users{}, fmt.Errorf("failed update data User: %v", err.Error())
	}
	return user, nil
}

func NewUsersUseCase(repo repository.UsersRepository) UsersUseCase {
	return &usersUseCase{repo: repo}
}
