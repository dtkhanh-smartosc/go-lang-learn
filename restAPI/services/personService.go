package services

import (
	"fmt"
	"net/http"
	"restAPI/model"

	"github.com/gin-gonic/gin"
)

type NewPerson struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

type PersonUpdate struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetPeople(c *gin.Context) {
	var people []model.Person
	db, err := model.Database()

	if err != nil {
		fmt.Println(err)
	}

	if err := db.Find(&people).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, people)
}
func GetPerson(c *gin.Context) {
	id := c.Param("id")
	var person model.Person
	db, err := model.Database()

	if err != nil {
		fmt.Println(err)
	}

	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "person not found"})
		return
	}

	c.JSON(http.StatusOK, person)
}

func PostPerson(c *gin.Context) {
	var person NewPerson

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPerson := model.Person{Name: person.Name, Age: person.Age}

	db, err := model.Database()
	if err != nil {
		fmt.Println(err)
	}

	if err := db.Create(&newPerson).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newPerson)
}
func DelPerson(c *gin.Context) {
	var person model.Person
	id := c.Param("id")
	db, err := model.Database()

	if err != nil {
		fmt.Println(err)
	}

	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "person not found"})
		return
	}

	if err := db.Delete(&person).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, person)

}
func UpdatePerson(c *gin.Context) {
	var person model.Person
	id := c.Param("id")
	db, err := model.Database()

	if err != nil {
		fmt.Println(err)
	}

	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "person not found"})
		return
	}

	var personUpdate PersonUpdate

	if err := c.ShouldBindJSON(&personUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println("personUpdate", personUpdate)

	if err := db.Model(&person).Updates(model.Person{Name: personUpdate.Name, Age: personUpdate.Age}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, person)
}
