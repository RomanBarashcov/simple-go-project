package repositories_test

import (
	utils "simple-go-project/src/tests/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func StartUserRepositoryTesting() {
	utils.WriteConsoleLog("Start UserRepository testring")
}

func TestFindUserById(t *testing.T) {

	StartUserRepositoryTesting()

	repository := &MUserRepository{}
	user := mockUsers[1]

	actual := repository.FindUserById(user.ID)
	expected := user

	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Email, actual.Email)

	EndUserRepositoryTesting()

}

func EndUserRepositoryTesting() {
	utils.WriteConsoleLog("End UserRepository testring")
}
