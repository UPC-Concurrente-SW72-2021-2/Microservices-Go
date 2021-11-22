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
	fmt.Fprintf(w, "Welcome the my GO API!")
}

func main() {
	// cnnnetwork.NeuronalNetworkCNN()

	repository.AddPerson(models.Person{
		IdPersona:          1,
		IdVacunadosCovid:   1,
		IdEess:             1,
		IdCentroVacunacion: 1,
		IdVacuna:           1,
		IdGrupoRiesgo:      1,
		Dosis:              0,
		Edad:               15,
	})
	repository.AddPerson(models.Person{
		IdPersona:          2,
		IdVacunadosCovid:   2,
		IdEess:             2,
		IdCentroVacunacion: 2,
		IdVacuna:           2,
		IdGrupoRiesgo:      2,
		Dosis:              1,
		Edad:               15,
	})

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/persons", handlers.AddPerson).Methods("POST")
	router.HandleFunc("/persons", handlers.GetPersons).Methods("GET")
	router.HandleFunc("/persons/{edad}", handlers.GetPersonsByAge).Methods("GET")
	// router.HandleFunc("/tasks/{id}", DetelePerson).Methods("DELETE")
	// router.HandleFunc("/tasks/{id}", UpdatePerson).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))
}
