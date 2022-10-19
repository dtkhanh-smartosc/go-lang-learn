package repository

import (
	"log"
	"restAPI/database"
	"restAPI/model"

	"gorm.io/gorm"
)

type PersonRepository struct {
	db *gorm.DB
}

func NewPersonRepository() *PersonRepository {
	return &PersonRepository{db: database.Connection}
}

func (repo *PersonRepository) CreatePerson(person *model.Person) error {
	log.Println(person)

	if err := repo.db.Create(person).Error; err != nil {
		log.Println("---------------------------")
		log.Println(err)
		return err
	}

	return nil
}
