package main

import (
	"github.com/DaviAlencar2/learning-go/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	router.Run()
}