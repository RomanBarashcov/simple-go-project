package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUserById(t *testing.T) {

	repository := &mockUserRepository{}
	user := mockUsers[1]

	actual := repository.FindUserById(user.ID)
	expected := user

	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Email, actual.Email)

}
