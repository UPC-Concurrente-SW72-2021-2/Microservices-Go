package repository

import (
	"tfC19Helper/svc/cnnnetwork"
	"tfC19Helper/svc/models"
)

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

func GetPersonsByAge(personAge int) ([]models.Person, int) {
	var personsAge []models.Person
	prediction := cnnnetwork.NeuronalNetworkCNN()
	for _, person := range persons {
		if person.Edad == personAge && person.Dosis == prediction {
			personsAge = append(personsAge, person)
		}
	}
	return personsAge, prediction
}
