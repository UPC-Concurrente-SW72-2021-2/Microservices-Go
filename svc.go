package main

import (
	"fmt"
	"log"
	"net/http"
	"tfC19Helper/svc/handlers"
	"tfC19Helper/svc/models"
	"tfC19Helper/svc/repository"

	"github.com/gorilla/mux"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wecome the my GO API!")
}

func main() {
	repository.AddPerson(models.Person{
		IdPersona:          1,
		IdVacunadosCovid:   1,
		IdEess:             1,
		IdCentroVacunacion: 1,
		IdVacuna:           1,
		IdGrupoRiesgo:      1,
		Dosis:              1,
		Edad:               1,
	})
	repository.AddPerson(models.Person{
		IdPersona:          2,
		IdVacunadosCovid:   2,
		IdEess:             2,
		IdCentroVacunacion: 2,
		IdVacuna:           2,
		IdGrupoRiesgo:      2,
		Dosis:              2,
		Edad:               2,
	})

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/persons", handlers.AddPerson).Methods("POST")
	router.HandleFunc("/persons", handlers.GetPersons).Methods("GET")
	router.HandleFunc("/persons/{id}", handlers.GetPersonById).Methods("GET")
	// router.HandleFunc("/tasks/{id}", DetelePerson).Methods("DELETE")
	// router.HandleFunc("/tasks/{id}", UpdatePerson).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))
}
