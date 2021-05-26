package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gunererd/dummy-challange/src/services"
	"gitlab.com/gunererd/dummy-challange/src/utils"
)

type Controller interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
	Info(c *gin.Context)
	ChangePassword(c *gin.Context)
	ListUsers(c *gin.Context)
}

type controller struct {
	userService  services.UserService
	tokenService services.TokenService
}

func NewController(userService services.UserService, tokenService services.TokenService) *controller {
	return &controller{
		userService:  userService,
		tokenService: tokenService,
	}
}

func (ctr controller) Signup(c *gin.Context) {
	type RequestUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var ru RequestUser
	if err := c.ShouldBindJSON(&ru); err != nil {
		errCode := http.StatusBadRequest
		err = &utils.FError{
			Code:      errCode,
			ErrorCode: "errors.badRequest",
			Message:   err.Error(),
		}

		c.JSON(errCode, err)
		return
	}

	ctr.userService.Create(ru.Username, ru.Password)
	c.JSON(http.StatusOK, gin.H{})
	return
}

func (ctr controller) Login(c *gin.Context) {
	type RequestUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var ru RequestUser
	if err := c.ShouldBindJSON(&ru); err != nil {

		errCode := http.StatusBadRequest
		err = &utils.FError{
			Code:      errCode,
			ErrorCode: "errors.badRequest",
			Message:   err.Error(),
		}

		c.JSON(errCode, err)
		return
	}

	u, err := ctr.userService.Get(ru.Username)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}

	if u.Password != ru.Password {

		errCode := http.StatusForbidden

		err := &utils.FError{
			Code:      errCode,
			ErrorCode: "errors.passwordMismatched",
			Message:   fmt.Sprintf("Password mismatched!"),
		}

		c.JSON(errCode, err)
		return
	}

	token, err := ctr.tokenService.GenerateToken(ru.Username)

	if err != nil {
		errCode := err.(*utils.FError).Code
		c.JSON(errCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
	return
}

func (ctr controller) Info(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")

	username, err := ctr.tokenService.ValidateToken(authHeader)

	if err != nil {
		errCode := err.(*utils.FError).Code
		c.JSON(errCode, err)
		return
	}

	user, err := ctr.userService.Get(username)
	if err != nil {
		errCode := err.(*utils.FError).Code
		c.JSON(errCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": user.Username})
	return
}

func (ctr controller) ChangePassword(c *gin.Context) {

	type RequestBody struct {
		Password string `json:"password"`
	}

	var ru RequestBody
	if err := c.ShouldBindJSON(&ru); err != nil {
		errCode := http.StatusBadRequest
		err := &utils.FError{
			Code:      errCode,
			ErrorCode: "errors.badRequest",
			Message:   "Bad request!",
		}

		c.JSON(errCode, err)
		return
	}

	authHeader := c.GetHeader("Authorization")
	username, err := ctr.tokenService.ValidateToken(authHeader)

	if err != nil {
		errCode := err.(*utils.FError).Code
		c.JSON(errCode, err)
		return
	}

	err = ctr.userService.ChangePassword(username, ru.Password)
	if err != nil {
		errCode := err.(*utils.FError).Code
		c.JSON(errCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "password changed!"})
	return
}

func (ctr controller) ListUsers(c *gin.Context) {

	users := ctr.userService.List()

	c.JSON(http.StatusOK, gin.H{"users": users})
	return
}
