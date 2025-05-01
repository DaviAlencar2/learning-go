package user

import (
    "github.com/DaviAlencar2/learning-go/internal/user/handlers"
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
    r.GET("/user/:userId", handlers.FindUserById)
    r.POST("/user/", handlers.CreateUser)
}