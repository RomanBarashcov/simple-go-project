package services

import "simple-go-project/src/repositories"

var bookRepository repositories.IBookRepository
var reviewRepository repositories.IReviewRepository
var userRepository repositories.IUserRepository

func Setup(bRep repositories.IBookRepository, rRep repositories.IReviewRepository, uRep repositories.IUserRepository) {

	bookRepository = bRep
	reviewRepository = rRep
	userRepository = uRep
}
