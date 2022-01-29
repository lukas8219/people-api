package service

import (
	"go-crud/data"

	"github.com/jinzhu/gorm"
)

type PersonService struct {
	Database *gorm.DB
}

func (p PersonService) CreatePerson(person *data.Person) (data.Person, error) {
	p.Database.Create(person)
	return *person, nil
}

func (p PersonService) UpdatePerson(id int, person *data.Person) (data.Person, error) {
	person.ID = id
	p.Database.Save(&person)
	return *person, nil
}

func (p PersonService) GetPerson(id int) (data.Person, error) {
	var person data.Person
	p.Database.Raw("SELECT * FROM go.people WHERE id = ?", id).Scan(&person)

	return person, nil
}
