package Controllers

import (
	"BackandContrat/Data/Service"
	"github.com/gin-gonic/gin"
)

func ListAllUsers(c *gin.Context) {
	repository := Service.UserRepository{}
	users := repository.ListAllUsers()
	c.JSON(200, users)
}
