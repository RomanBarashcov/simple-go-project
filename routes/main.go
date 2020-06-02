package routes

import "code/services"

var bookService services.IBookService
var reviewService services.IReviewService

func Setup(bSer services.IBookService, rSer services.IReviewService) {
	bookService = bSer
	reviewService = rSer
}
