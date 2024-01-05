package usecase

import (
	"field_work/entity/dto"
	"field_work/shared/service"
	
)

type AuthUseCase interface {
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
}

type authUseCase struct {
	userUC     UsersUseCase
	jwtService service.JwtService
}

// Login implements AuthUseCase.
func (a *authUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	user, err := a.userUC.FindUsersForLogin(payload.User, payload.Password)
	if err != nil{
		return dto.AuthResponseDto{}, err
	}
	token, err := a.jwtService.CreateToken(user)
	if err != nil{
		return dto.AuthResponseDto{}, err
	}
	return token, nil
}

func NewAuthUseCase(userUC UsersUseCase, jwtService service.JwtService) AuthUseCase {
	return &authUseCase{userUC: userUC, jwtService: jwtService}

}
