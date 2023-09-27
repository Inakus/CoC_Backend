package main

import (
	"github.com/Inakus/CoC_Backend/controllers"
	"github.com/Inakus/CoC_Backend/initializers"
	"github.com/Inakus/CoC_Backend/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.UserAuth, controllers.Validate)
	r.GET("/logout", middleware.UserAuth, controllers.Logout)
	r.GET("/admin", middleware.AdminAuth, controllers.Validate)

	r.Run()
}
