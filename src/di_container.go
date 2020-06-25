package main

import (
	"simple-go-project/src/repositories"
	"simple-go-project/src/routes"
	"simple-go-project/src/services"
)

func MakeInjection() {

	InjectRepositories()

}

func InjectRepositories() {

	var bookRep repositories.IBookRepository = &repositories.BookRepository{}
	var reviewRep repositories.IReviewRepository = &repositories.ReviewRepository{}
	var userRep repositories.IUserRepository = &repositories.UserRepository{}

	InjectService(bookRep, reviewRep, userRep)
}

func InjectService(bRep repositories.IBookRepository, rRep repositories.IReviewRepository, uRep repositories.IUserRepository) {

	services.Setup(bRep, rRep, uRep)

	var bookSer services.IBookService = &services.BookService{}
	var reviewSer services.IReviewService = &services.ReviewService{}

	InjectRoutes(bookSer, reviewSer)

}

func InjectRoutes(bSer services.IBookService, rSer services.IReviewService) {

	routes.Setup(bSer, rSer)

}
