package repository

import "tfC19Helper/svc/models"

var persons []models.Person
var nextID = 1

func GetPersons() []models.Person {
	return persons
}

func AddPerson(person models.Person) int {
	person.ID = nextID
	nextID++
	persons = append(persons, person)
	return person.ID
}
