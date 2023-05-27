package User

import (
	"go-ddd-rest-api-sample/src/Infrastructures"
	"go-ddd-rest-api-sample/src/Infrastructures/Repositories"
	"go-ddd-rest-api-sample/src/UseCases/User"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct{}

func (controller *UserController) GetUser(c echo.Context) error {
	userIdInt, _ := strconv.Atoi(c.Param("userId"))

	repository := Repositories.NewUserRepository(Infrastructures.GetDb())
	useCase := User.NewGetUserUseCase(repository)

	userDto := useCase.GetUser(userIdInt)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": map[string]interface{}{
			"userId":    userDto.Id,
			"firstName": userDto.FirstName,
			"lastName":  userDto.LastName,
		},
	})
}
