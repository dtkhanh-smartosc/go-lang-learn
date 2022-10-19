package controller

import (
	"gormGinAPI/model"
	"gormGinAPI/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePerson(c *gin.Context) {
	var person model.Person

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	personServices := services.NewPersonServices()
	if err := personServices.CreatePerson(&person); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, person)

}
