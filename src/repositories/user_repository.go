package repositories

import (
	"simple-go-project/src/database/config"
	"simple-go-project/src/entities"
)

type IUserRepository interface {
	FindUserById(id int64) *entities.User
}

type UserRepository struct{}

func (u *UserRepository) FindUserById(id int64) *entities.User {

	db := config.GetConnection()
	defer db.Close()

	user := new(entities.User)
	db.Table("users as u").Select("u.*").Where("u.ID = ?", id).Scan(&user)

	return user
}
