package main

import (
	"gormGinAPI/controller"
	"gormGinAPI/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	connection, err := database.ConnectSqlite()
	if err != nil {
		log.Fatal(err)
	}

	database.Connection = connection
	router := gin.Default()

	router.POST("/people", controller.CreatePerson)

	router.Run("localhost:5050")
}
