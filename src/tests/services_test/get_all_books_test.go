package services_test

import (
	"testing"

	r "simple-go-project/src/repositories"

	. "../repositories_test"

	"github.com/stretchr/testify/assert"
)

var (
	bRep r.IBookRepository
	rRep r.IReviewRepository
	uRep r.IUserRepository
)

func TestGetAllBooks(t *testing.T) {
	assert.Equal(t, true, true)
}

func SetupBookService() {

	bReprr := &MBookRepository{}

}
