package repository

import (
	"gormGinAPI/database"
	"gormGinAPI/model"
	"log"

	"gorm.io/gorm"
)

type PersonRepository struct {
	db *gorm.DB
}

func NewPersonRepository() *PersonRepository {
	return &PersonRepository{db: database.Connection}
}

func (personRepo *PersonRepository) CreatePerson(person *model.Person) error {
	log.Println(person)
	log.Println("personRepo", personRepo)
	result := personRepo.db.Create(person)
	//testPerson := model.Person{Name: "khanhdt", Age: 12}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
