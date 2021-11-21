package models

type Person struct {
	ID                 int `json:"id,omitempty"`
	IdVacunadosCovid   int `json:"vacunadosCovid"`
	IdEess             int `json:"eess"`
	IdCentroVacunacion int `json:"centroVacunacion"`
	IdVacuna           int `json:"vacuna"`
	IdGrupoRiesgo      int `json:"grupoRiesgo"`
	Dosis              int `json:"dosis"`
	Edad               int `json:"edad"`
}
