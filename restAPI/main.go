package main

import (
	"restAPI/controller"
	"restAPI/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectSqlLite()

	router := gin.Default()
	router.POST("/people", controller.CreatePerson)

	// router.GET("/people", services.GetPeople)
	// router.GET("/people/:id", services.GetPerson)
	// router.POST("/people", services.NewPersonService().CreatePerson)
	// router.DELETE("/people/:id", services.DelPerson)
	// router.PUT("/people/:id", services.UpdatePerson)
	router.Run("localhost:5050")

}
