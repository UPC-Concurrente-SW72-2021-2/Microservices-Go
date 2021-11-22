package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"tfC19Helper/svc/models"
	"tfC19Helper/svc/repository"

	"github.com/gorilla/mux"
)

// func HandleRequest(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Request received: ", r.Method)

// 	switch r.Method {
// 	case http.MethodGet:
// 		list(w, r)
// 		break
// 	// case http.MethodPost:
// 	// 	add(w, r)
// 	// 	break
// 	default:
// 		w.WriteHeader(405)
// 		w.Write([]byte("Method not allowed"))
// 		break
// 	}
// }

// func list(w http.ResponseWriter, r *http.Request) {
// 	persons := repository.GetPersons()
// 	json, _ := json.Marshal(persons)

// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(200)
// 	w.Write(json)

// 	log.Println("Response Returned: ", 200)
// }

// func add(w http.ResponseWriter, r *http.Request) {
// 	payload, _ := ioutil.ReadAll(r.Body)

// 	var person models.Person
// 	err := json.Unmarshal(payload, &person)

// 	if err != nil || person.IdVacunadosCovid == 0 || person.IdEess == 0 || person.IdCentroVacunacion == 0 || person.IdVacuna == 0 || person.IdGrupoRiesgo == 0 || person.Dosis == 0 || person.Edad == 0 {
// 		w.WriteHeader(400)
// 		w.Write([]byte("Bad Request"))
// 		return
// 	}

// 	person.ID = repository.AddPerson(person)

// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(201)

// 	json, _ := json.Marshal(person)
// 	w.Write(json)

// 	log.Println("Response Returned: ", 201)
// }

// func HandlePredict(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Request received: ", r.Method)
// // 	persons :=repository.GetPersons()
// // 	edad:=r.FormValue("edad")

// // 	w.Header().Add("Content-Type", "application/json")
// }

func AddPerson(w http.ResponseWriter, r *http.Request) {
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

func GetPersons(w http.ResponseWriter, r *http.Request) {
	persons := repository.GetPersons()
	json, _ := json.Marshal(persons)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)

	log.Println("Response Returned: ", 200)
}

func GetPersonByAge(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personID, err := strconv.Atoi(vars["edad"])
	person := repository.GetPersonById(personID)

	if err != nil {
		return
	}

	if person.ID == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}

	json, _ := json.Marshal(person)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}
