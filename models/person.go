package models

type Person struct {
	ID                 int `json:"id,omitempty"`
	IdPersona          int `json:"persona"`
	IdVacunadosCovid   int `json:"vacunadosCovid"`
	IdEess             int `json:"eess"`
	IdCentroVacunacion int `json:"centroVacunacion"`
	IdVacuna           int `json:"vacuna"`
	IdGrupoRiesgo      int `json:"grupoRiesgo"`
	Dosis              int `json:"dosis"`
	Edad               int `json:"edad"`
}
