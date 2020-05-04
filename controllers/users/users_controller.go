package handler_users

import (
	"github.com/gabrielnotong/bookstore_users-api/domain/users"
	"github.com/gabrielnotong/bookstore_users-api/errors"
	"github.com/gabrielnotong/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	// strconv.ParseInt parameters
	base10 = 10
	bit64  = 64
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, err := strconv.ParseInt(userIdParam, base10, bit64)
	if err != nil {
		restErr := errors.NewBadRequestError("User id should be a number")
		return 0, restErr
	}

	return userId, nil
}

func Create(c *gin.Context) {
	var u users.User

	if err := c.ShouldBindJSON(&u); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, restErr := services.UsersService.CreateUser(u)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func FindUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, restErr := services.UsersService.FindUser(userId)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	var u users.User

	userId, idErr := getUserId(c.Param("id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := c.ShouldBindJSON(&u); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	u.Id = userId
	cu, updateErr := services.UsersService.UpdateUser(&u)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, cu.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if delErr := services.UsersService.DeleteUser(userId); delErr != nil {
		c.JSON(delErr.Status, delErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	uu, findErr := services.UsersService.Search(status)
	if findErr != nil {
		c.JSON(findErr.Status, findErr)
		return
	}

	c.JSON(http.StatusOK, uu.Marshall(c.GetHeader("X-Public") == "true"))
}

func Login(c *gin.Context) {
	lr := users.LoginRequest{}

	if err := c.ShouldBindJSON(&lr); err != nil {
		bErr := errors.NewBadRequestError(err.Error())
		c.JSON(bErr.Status, bErr)
		return
	}

	u, err := services.UsersService.Login(lr)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, u.Marshall(c.GetHeader("X-Public") == "true"))
}