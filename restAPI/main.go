package main

import (
	"fmt"
	"restAPI/model"
	"restAPI/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := model.Database()

	if err != nil {
		fmt.Println(err)
	}

	db.DB()

	router := gin.Default()

	router.GET("/people", services.GetPeople)
	router.GET("/people/:id", services.GetPerson)
	router.POST("/people", services.PostPerson)
	router.DELETE("/people/:id", services.DelPerson)
	router.PUT("/people/:id", services.UpdatePerson)
	router.Run("localhost:5050")

}
