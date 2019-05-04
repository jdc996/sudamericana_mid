package models

type Ciudad struct {
	Id     int
	Nombre string
	PaisId *Pais
}
