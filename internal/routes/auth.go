package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jz222/loggy/internal/controllers"
	"github.com/jz222/loggy/internal/store"
)

func authRoutes(router *gin.RouterGroup, store store.InterfaceStore) {
	controller := controllers.GetAuthControllers(store)

	router.POST("/setup", controller.Setup)
	router.POST("/signup", controller.SignUp)
	router.POST("/signin", controller.SignIn)
}
