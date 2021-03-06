package tests

import (
	"simple-go-project/src/entities"
	"simple-go-project/src/repositories"
)

type MUserRepository struct{}

var (
	MockUserRepository repositories.IUserRepository = &MUserRepository{}
)

var mockUsers = []entities.User{
	entities.User{
		ID:    1,
		Name:  "Some Name1",
		Email: "Some Email1",
	},
	entities.User{
		ID:    2,
		Name:  "Some Name2",
		Email: "Some Email2",
	},
	entities.User{
		ID:    3,
		Name:  "Some Name3",
		Email: "Some Email3",
	},
}

func (m *MUserRepository) FindUserById(id int64) *entities.User {

	for i, v := range mockUsers {
		if v.ID == id {

			return &mockUsers[i]

		}
	}

	return nil

}
