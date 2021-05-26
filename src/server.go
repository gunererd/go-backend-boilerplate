package src

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gunererd/dummy-challange/src/controllers"
	"gitlab.com/gunererd/dummy-challange/src/repositories"
	"gitlab.com/gunererd/dummy-challange/src/services"
)

func NewServer() *gin.Engine {

	userStore := repositories.NewUserStore()

	userService := services.NewUserService(userStore)
	tokenService := services.NewTokenService()

	controller := controllers.NewController(userService, tokenService)

	server := gin.New()

	server.POST("/login/", controller.Login)
	server.POST("/change-password/", controller.ChangePassword)
	server.GET("/info/", controller.Info)
	server.POST("/signup/", controller.Signup)
	server.GET("/users/", controller.ListUsers)

	return server
}
