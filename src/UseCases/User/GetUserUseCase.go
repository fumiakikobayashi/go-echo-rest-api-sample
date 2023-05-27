package User

import (
	"go-ddd-rest-api-sample/src/Domains"
	"go-ddd-rest-api-sample/src/Infrastructures/Repositories"
	"go-ddd-rest-api-sample/src/UseCases/Dto"
)

type GetUserUseCaseInterface interface {
	GetUser(userIdInt int) Dto.UserDto
}

type getUserUseCase struct {
	repository Repositories.UserRepositoryInterface
}

func NewGetUserUseCase(repository Repositories.UserRepositoryInterface) GetUserUseCaseInterface {
	return &getUserUseCase{repository}
}

func (useCase *getUserUseCase) GetUser(userIdInt int) Dto.UserDto {
	userId := Domains.UserId{
		Value: userIdInt,
	}
	user, err := useCase.repository.FindById(userId)
	if err != nil {
		return Dto.UserDto{}
	}

	return Dto.UserDto{
		Id:        userId.GetValue(),
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
	}
}
