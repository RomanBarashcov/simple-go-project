package repositories_test

import (
	"simple-go-project/src/entities"
	"simple-go-project/src/repositories"
)

type MReviewRepository struct{}

var (
	MockReviewRepository repositories.IReviewRepository = &MReviewRepository{}
)

var mockReviews = []entities.Review{
	entities.Review{
		BookID: 1,
		Rating: 5,
		Text:   "Best book I have ever read1",
		UserID: 1,
	},
	entities.Review{
		BookID: 2,
		Rating: 5,
		Text:   "Best book I have ever read2",
		UserID: 2,
	},
	entities.Review{
		BookID: 3,
		Rating: 5,
		Text:   "Best book I have ever read3",
		UserID: 3,
	},
}

func (r *MReviewRepository) CreateReview(newReview *entities.Review) *entities.Review {

	newReview.ID = int64(len(mockReviews)) + 1
	mockReviews = append(mockReviews, *newReview)
	return newReview

}

func (r *MReviewRepository) UpdateReview(upReview *entities.Review) *entities.Review {

	for i, v := range mockReviews {

		if v.ID == upReview.ID {
			mockReviews[i].BookID = upReview.BookID
			mockReviews[i].Rating = upReview.Rating
			mockReviews[i].Text = upReview.Text
			mockReviews[i].UserID = upReview.UserID
		}

	}

	return upReview

}

func (r *MReviewRepository) DeleteReview(id int64) bool {

	for i, v := range mockReviews {

		if v.ID == id {

			mockReviews = append(mockReviews[:i], mockReviews[i+1:]...) // delete element by index
			return true
		}

	}

	return false

}
