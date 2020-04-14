package handler_users

import (
	"github.com/gabrielnotong/bookstore_users-api/domain/users"
	"github.com/gabrielnotong/bookstore_users-api/errors"
	"github.com/gabrielnotong/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	var u users.User

	if err := c.ShouldBindJSON(&u); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, restErr := services.CreateUser(u)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Search(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("User id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, restErr := services.FindUser(userId)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, user)
}
