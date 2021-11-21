package main

import (
	"log"
	"net/http"
	"tfC19Helper/svc/handlers"
	"tfC19Helper/svc/models"
	"tfC19Helper/svc/repository"
)

func main() {
	repository.AddPerson(models.Person{
		IdVacunadosCovid:   1,
		IdEess:             1,
		IdCentroVacunacion: 1,
		IdVacuna:           1,
		IdGrupoRiesgo:      1,
		Dosis:              1,
		Edad:               1,
	})
	repository.AddPerson(models.Person{
		IdVacunadosCovid:   2,
		IdEess:             2,
		IdCentroVacunacion: 2,
		IdVacuna:           2,
		IdGrupoRiesgo:      2,
		Dosis:              2,
		Edad:               2,
	})

	http.HandleFunc("/", handlers.HandleRequest)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
