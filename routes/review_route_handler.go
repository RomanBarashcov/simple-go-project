package routes

import (
	"code/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateReview(c *gin.Context) {

	newReview := &entities.Review{}
	c.Bind(newReview)

	review := reviewService.CreateReview(newReview)
	c.JSON(200, review)

}

func UpdateReview(c *gin.Context) {

	upReview := &entities.Review{}
	c.Bind(upReview)

	review := reviewService.UpdateReview(upReview)
	c.JSON(200, review)

}

func DeleteReview(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, "Payload is incorrect!")
	}

	success := reviewService.DeleteReview(id)

	c.PureJSON(200, gin.H{
		"deleted": success,
	})

}
