package services

import (
	"code/entities"
)

type IReviewService interface {
	CreateReview(newReview *entities.Review) *entities.Review
	UpdateReview(upReview *entities.Review) *entities.Review
	DeleteReview(id int64) bool
}

type ReviewService struct{}

func (r ReviewService) CreateReview(newReview *entities.Review) *entities.Review {

	book := bookRepository.FindBookById(newReview.BookID)
	if book == nil {
		panic("Book does not exists!")
	}

	user := userRepository.FindUserById(newReview.UserID)
	if user == nil {
		panic("User does not exists!")
	}

	review := reviewRepository.CreateReview(newReview)
	if review == nil {
		panic("Review does not exists!")
	}

	review.UserEmail = user.Email
	review.UserName = user.Name

	return review

}

func (r ReviewService) UpdateReview(upReview *entities.Review) *entities.Review {

	book := bookRepository.FindBookById(upReview.BookID)
	if book == nil {
		panic("Book does not exists!")
	}

	user := userRepository.FindUserById(upReview.UserID)
	if user == nil {
		panic("User does not exists!")
	}

	review := reviewRepository.UpdateReview(upReview)
	if review == nil {
		panic("Review does not exists!")
	}

	review.UserEmail = user.Email
	review.UserName = user.Name

	return review

}

func (r ReviewService) DeleteReview(id int64) bool {

	success := reviewRepository.DeleteReview(id)
	return success
}
