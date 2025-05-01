package handlers

import (
	"fmt"

	"github.com/DaviAlencar2/learning-go/internal/config"
	"github.com/DaviAlencar2/learning-go/internal/user/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := config.NewBadRequestErr(
			fmt.Sprintf("Erro em algum campo, error=%s",err.Error()),
		)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)

}