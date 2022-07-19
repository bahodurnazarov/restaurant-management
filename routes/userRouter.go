package routes

import (
	controller "github.com/bahodurnazarov/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.Get("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("users/signup", controller.SingUp())
	incomingRoutes.POST("users/login", controller.Login())
	
}