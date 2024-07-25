package Routes

import (
	"BackandContrat/Application/Controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(routGroup *gin.RouterGroup) {
	apiRoutes := routGroup.Group("/api")
	{
		userRoutes := apiRoutes.Group("/users")
		{
			userRoutes.GET("/all-users", Controllers.ListAllUsers)
			createUsers := userRoutes.Group("/create")
			{
				createUsers.POST("/recruiter", Controllers.CreateRecrutadorController)
				createUsers.POST("/candidate", Controllers.CreateCandidatoController)
			}
		}
	}
}
