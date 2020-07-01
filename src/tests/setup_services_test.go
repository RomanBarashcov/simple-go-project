package tests

import (
	r "simple-go-project/src/repositories"
	s "simple-go-project/src/services"
)

var (
	bRep r.IBookRepository
	rRep r.IReviewRepository
	uRep r.IUserRepository

	bSer s.IBookService
	rSer s.IReviewService
)

func SetupServiceIngection() {

	bRep = &MBookRepository{}
	rRep = &MReviewRepository{}
	uRep = &MUserRepository{}

	s.Setup(bRep, rRep, uRep)

	bSer = &s.BookService{}
	rSer = &s.ReviewService{}

}
