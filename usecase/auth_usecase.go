package usecase

import (
	"field_work/entity"
	"field_work/entity/dto"
	"field_work/shared/service"
	"strings"
)

type AuthUseCase interface {
	Register(payload entity.Users, role string) (entity.Users, error)
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
}

type authUseCase struct {
	userUC     UsersUseCase
	jwtService service.JwtService
}

func (a *authUseCase) Register(payload entity.Users, role string) (entity.Users, error) {
	if strings.Contains(role, "admin") {
		payload.Role = "Admin"
	} else {
		payload.Role = "Customer"
	}

	user, err := a.userUC.RegisterNewUsers(payload)
	return user, err
}

func (a *authUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	user, err := a.userUC.FindUsersForLogin(payload.User, payload.Password)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	token, err := a.jwtService.CreateToken(user)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	return token, nil
}

func NewAuthUseCase(userUC UsersUseCase, jwtService service.JwtService) AuthUseCase {
	return &authUseCase{userUC: userUC, jwtService: jwtService}

}
