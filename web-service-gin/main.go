package main

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var people = []person{
	{ID: "1", Name: "Adt", Age: 21},
	{ID: "2", Name: "Bdt", Age: 22},
	{ID: "3", Name: "Cdt", Age: 23},
}

/*
	gin.Context It carries request details, validates and serializes JSON, and more.
 	(Despite the similar name, this is different from
 	Go’s built-in context package.)
*/

// get people
func getPeople(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, people)
}

// get person
func getPerson(c *gin.Context) {
	id := c.Param("id")
	for _, p := range people {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}

// post new person
func postPerson(c *gin.Context) {
	var newPerson person

	if err := c.BindJSON(&newPerson); err != nil {
		return
	}
	people = append(people, newPerson)
	c.IndentedJSON(http.StatusCreated, people)
}

// remove person
func removePerson(c *gin.Context) {
	var id = c.Param("id")

	for i, p := range people {
		if p.ID == id {
			// only pass element but not the whole slice
			people = append(people[:i], people[i+1:]...)
			c.IndentedJSON(http.StatusOK, people)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}

// put-update person
func upadtePerson(c *gin.Context) {
	var id = c.Param("id")
	var personInfo person
	if err := c.BindJSON(&personInfo); err != nil {
		return
	}
	val := reflect.ValueOf(personInfo)
	types := val.Type()
	for index, p := range people {
		if p.ID == id { // finding person with the same id
			for i := 0; i < val.NumField(); i++ {
				if !val.Field(i).IsZero() { // check if the value is exits
					switch types.Field(i).Name {
					case "Name":
						p.Name = val.Field(i).String()
						people[index] = p
					case "Age":
						p.Age = int(val.Field(i).Int())
						people[index] = p
					}
				}
			}
			c.IndentedJSON(http.StatusOK, people)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})

}

func main() {
	router := gin.Default()

	router.GET("/people", getPeople)
	router.GET("people/:id", getPerson)
	router.POST("/people", postPerson)
	router.PUT("people/:id", upadtePerson)
	router.DELETE("people/:id", removePerson)

	router.Run("localhost:5050")
}
