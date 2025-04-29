package routes

import (
	"github.com/DaviAlencar2/learning-go/internal/user/handlers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup){
	r.GET("/user/:userId", handlers.FindUserById)
	r.POST("/user/")
	r.PUT("/user/:userId")
	r.DELETE("/user/:UserId")
}