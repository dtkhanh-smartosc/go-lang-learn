package services

import (
	"gormGinAPI/model"
	"gormGinAPI/repository"
)

type PersonService struct {
	personRepository *repository.PersonRepository
}

func NewPersonServices() *PersonService {
	return &PersonService{personRepository: repository.NewPersonRepository()}
}

func (personService *PersonService) CreatePerson(person *model.Person) error {
	if err := personService.personRepository.CreatePerson(person); err != nil {
		return err
	}
	return nil
}
