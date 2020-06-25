package repositories

import (
	"simple-go-project/src/database/config"
	"simple-go-project/src/database/models"
	"simple-go-project/src/entities"
)

type IReviewRepository interface {
	CreateReview(newReview *entities.Review) *entities.Review
	UpdateReview(upReview *entities.Review) *entities.Review
	DeleteReview(id int64) bool
}

type ReviewRepository struct{}

func (r *ReviewRepository) CreateReview(newReview *entities.Review) *entities.Review {

	db := config.GetConnection()
	defer db.Close()

	review := models.Review{
		Text:   newReview.Text,
		Rating: newReview.Rating,
		BookID: newReview.BookID,
		UserID: newReview.UserID,
	}

	newRecordResult := db.NewRecord(review)
	if !newRecordResult {
		panic("new Record book error")
	}

	db.Create(&review).Scan(&newReview)
	return newReview

}

func (r *ReviewRepository) UpdateReview(upReview *entities.Review) *entities.Review {

	db := config.GetConnection()
	defer db.Close()

	review := models.Review{ID: upReview.ID}
	db.First(&review)

	review.BookID = upReview.BookID
	review.Rating = upReview.Rating
	review.Text = upReview.Text
	review.UserID = upReview.UserID

	upResult := db.Save(&review).Value
	if upResult == nil {
		panic("Error when update review")
	}

	return upReview

}

func (r *ReviewRepository) DeleteReview(id int64) bool {

	var success bool = false

	db := config.GetConnection()
	defer db.Close()

	review := models.Review{ID: id}
	result := db.Delete(&review).Value
	if result == nil {
		panic("Error when delete review")
	}

	success = true
	return success

}
