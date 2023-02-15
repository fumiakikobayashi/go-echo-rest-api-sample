package Repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-ddd-rest-api-sample/src/Domains"
	"go-ddd-rest-api-sample/src/Infrastructures/Models"
)

type UserRepositoryInterface interface {
	FindById(id Domains.UserId) (*Domains.User, error)
}

type userRepository struct {
	db gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{db: *db}
}

func (repository *userRepository) FindById(userId Domains.UserId) (*Domains.User, error) {
	userModel := Models.User{}
	if repository.db.Where("id = ?", userId.GetValue()).Find(&userModel).RecordNotFound() {
		// 例外処理
		fmt.Println("取得に失敗しました")
	}

	return Domains.NewUser(
			Domains.UserId{Value: userModel.Id},
			userModel.FirstName,
			userModel.LastName,
		),
		nil
}
