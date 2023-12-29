package routes

import (
	controller "mongodbex/controllers"

	"github.com/gin-gonic/gin"
)

// UserRoutes function
func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}
