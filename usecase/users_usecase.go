package usecase

import (
	"errors"
	"field_work/entity"
	"field_work/entity/dto"
	"field_work/repository"
	"field_work/shared/model"
	"fmt"
	"reflect"
	"time"

	"github.com/jinzhu/copier"
)

type Model interface {
	dto.UpdateUserDTO |
		entity.Users
}

type UsersUseCase interface {
	FindUsersByID(id string) (entity.Users, error)
	FindUsersByUsername(username string) (entity.Users, error)
	FindUsersForLogin(username, password string) (entity.Users, error)
	RegisterNewUsers(payload entity.Users) (entity.Users, error)
	UpdateUsers(payload dto.UpdateUserDTO, id string) (entity.Users, error)
	DeleteUsers(id string) error
	ListAll(page, size int) ([]entity.Users, model.Paging, error)
}

type usersUseCase struct {
	repo repository.UsersRepository
}

func (u *usersUseCase) DeleteUsers(id string) error {
	if _, err := u.repo.GetUsersByID(id); err != nil {
		return err
	}

	return u.repo.DeleteUser(id)
}

func (u *usersUseCase) FindUsersByID(id string) (entity.Users, error) {
	if id == "" {
		return entity.Users{}, errors.New("id harus diisi")
	}
	return u.repo.GetUsersByID(id)
}

func (u *usersUseCase) FindUsersByUsername(username string) (entity.Users, error) {
	if username == "" {
		return entity.Users{}, errors.New("username harus diisi")
	}
	return u.repo.GetUsersByUsername(username)
}

func (u *usersUseCase) FindUsersForLogin(username string, password string) (entity.Users, error) {
// 	if username == "" || password == "" {
// 		return entity.Users{}, errors.New("gagal masuk")
// 	}
	return u.repo.GetUsersByUsernameForLogin(username, password)
}

func (u *usersUseCase) ListAll(page int, size int) ([]entity.Users, model.Paging, error) {
	users, paging, err := u.repo.List(page, size)
	for i, user := range users {
		createdAt, _ := time.Parse("2006-01-02T15:04:05+07:00", user.CreatedAt)
		users[i].CreatedAt = createdAt.Format("Monday 02, January 2006 15:04:05")
	}

	return users, paging, err
}

func (u *usersUseCase) RegisterNewUsers(payload entity.Users) (entity.Users, error) {
	user, err := u.repo.CreateUser(payload)
	if err != nil {
		return entity.Users{}, fmt.Errorf("failed save data User: %v", err.Error())
	}
	return user, nil
}

func (u *usersUseCase) UpdateUsers(payload dto.UpdateUserDTO, id string) (entity.Users, error) {
	user, err := u.repo.GetUsersByID(id)
	if err != nil {
		return entity.Users{}, fmt.Errorf(fmt.Sprintf("User with ID '%s' not found", id))
	}

	userDTOMap := structToMap[dto.UpdateUserDTO](payload)
	userMap := structToMap[entity.Users](user)

	// Mengisi nilai payload yang kosong dengan nilai yang sudah ada d database
	for key, v := range userDTOMap {
		// Jika field userDTOMap/payload kosong
		if *v == "" {
			// isi field dengan nilai usermap
			*v = *userMap[key]
		}
	}

	// Mengisi nilai payload dengan nilai userDTOMap
	for key, value := range userDTOMap {
		field := reflect.ValueOf(&payload).Elem().FieldByName(key)

		if field.Kind() == reflect.String {
			fieldValue := *value
			field.SetString(fieldValue)
		}
	}

	// Meng-copy nilai dari payload ke struct user menggunakan package github.com/jinzhu/copier
	if err := copier.Copy(&user, &payload); err != nil {
		return entity.Users{}, fmt.Errorf("failed to copy user struct: %v", err.Error())
	}

	newUser, err := u.repo.UpdateUser(user)
	if err != nil {
		return entity.Users{}, fmt.Errorf("failed update data User: %v", err.Error())
	}

	return newUser, nil
}

func NewUsersUseCase(repo repository.UsersRepository) UsersUseCase {
	return &usersUseCase{repo: repo}
}

func structToMap[M Model](s M) map[string]*string {
	structValue := reflect.ValueOf(s)
	structType := structValue.Type()

	result := make(map[string]*string)

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldName := structType.Field(i).Name
		fieldValue := field.String()
		result[fieldName] = &fieldValue
	}

	return result
}
