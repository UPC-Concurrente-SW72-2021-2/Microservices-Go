package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"tfC19Helper/svc/models"
	"tfC19Helper/svc/repository"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received: ", r.Method)

	switch r.Method {
	case http.MethodGet:
		list(w, r)
		break
	case http.MethodPost:
		add(w, r)
		break
	default:
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		break
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	persons := repository.GetPersons()
	json, _ := json.Marshal(persons)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)

	log.Println("Response Returned: ", 200)
}

func add(w http.ResponseWriter, r *http.Request) {
	payload, _ := ioutil.ReadAll(r.Body)

	var person models.Person
	err := json.Unmarshal(payload, &person)

	if err != nil || person.IdVacunadosCovid == 0 || person.IdEess == 0 || person.IdCentroVacunacion == 0 || person.IdVacuna == 0 || person.IdGrupoRiesgo == 0 || person.Dosis == 0 || person.Edad == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}

	person.ID = repository.AddPerson(person)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)

	json, _ := json.Marshal(person)
	w.Write(json)

	log.Println("Response Returned: ", 201)
}
