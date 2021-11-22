package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tfC19Helper/svc/handlers"
	"tfC19Helper/svc/models"
	"tfC19Helper/svc/repository"

	"github.com/gorilla/mux"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome the my GO API!")
}

func main() {
	var personDb []models.PersonJs

	data, _ := ioutil.ReadFile("data/DB_VACUNACION_COVID19.json")

	err := json.Unmarshal(data, &personDb)

	if err != nil {
		fmt.Println(err)
	}

	for i := range personDb {
		repository.AddPerson(models.Person{
			IdPersona:          personDb[i].IdPersona,
			IdVacunadosCovid:   personDb[i].IdVacunadosCovid,
			IdEess:             personDb[i].IdEess,
			IdCentroVacunacion: personDb[i].IdCentroVacunacion,
			IdVacuna:           personDb[i].IdVacuna,
			IdGrupoRiesgo:      personDb[i].IdGrupoRiesgo,
			Dosis:              personDb[i].Dosis,
			Edad:               personDb[i].Edad,
		})
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/persons", handlers.AddPerson).Methods("POST")
	router.HandleFunc("/persons", handlers.GetPersons).Methods("GET")
	router.HandleFunc("/persons/{edad}", handlers.GetPersonsByAge).Methods("GET")

	log.Fatal(http.ListenAndServe(":80", router))
}
