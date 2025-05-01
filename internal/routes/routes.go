package routes

import (
	"github.com/DaviAlencar2/learning-go/internal/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup){
	user.RegisterRoutes(r)
}