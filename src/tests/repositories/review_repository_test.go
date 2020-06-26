package repositories_test

import (
	"simple-go-project/src/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateReview(t *testing.T) {

	repository := &mockReviewRepository{}
	review := new(entities.Review)

	review.BookID = 1
	review.Rating = 5
	review.Text = "SomeText"
	review.UserID = 1

	actual := repository.CreateReview(review)
	review.ID = 4
	expected := review

	assert.Equal(t, expected.ID, actual.ID)
}

func TestUpdateReview(t *testing.T) {

	repository := &mockReviewRepository{}
	review := mockReviews[1]

	review.BookID = 3
	review.Rating = 5
	review.Text = "SomeText5"
	review.UserID = 1

	actual := repository.UpdateReview(&review)
	expected := review

	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Text, actual.Text)
}

func TestDeleteReview(t *testing.T) {

	repository := &mockReviewRepository{}
	review := mockReviews[1]

	actual := repository.DeleteReview(review.ID)
	expected := true

	assert.Equal(t, expected, actual)

}
