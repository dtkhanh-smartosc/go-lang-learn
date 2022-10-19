package controller

import (
	"net/http"
	"restAPI/model"
	"restAPI/services"

	"github.com/gin-gonic/gin"
)

func CreatePerson(c *gin.Context) {
	var person *model.Person

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	personService := services.NewPersonService()

	if err := personService.CreatePerson(person); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, person)
	return
}
