package Controllers

import (
	"BackandContrat/Application/Dto"
	"BackandContrat/Data/Service"
	"BackandContrat/Domain/Factory"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCandidatoController(c *gin.Context) {
	var req Dto.CreateUserRequestDto
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	factory := Factory.CandidatoFactory{}
	newCandidato := factory.Create(req.Name, req.Username, req.Password)
	newRepository := Service.UserRepository{}

	newRepository.CreateUser(&factory, req.Name, req.Username, req.Password)

	c.JSON(http.StatusOK, newCandidato)
}

func CreateRecrutadorController(c *gin.Context) {
	var req Dto.CreateUserRequestDto
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	factory := Factory.RecrutadorFactory{}
	newRecrutador := factory.Create(req.Name, req.Username, req.Password)
	newRepository := Service.UserRepository{}

	newRepository.CreateUser(&factory, req.Name, req.Username, req.Password)

	c.JSON(http.StatusOK, newRecrutador)
}
