package models

import "time"

type Partidos struct {
	Id                int
	Fecha             time.Time
	EquipoLocalId     int
	EquipoVisitanteId int
	GolesLocal        int
	GolesVisitante    int
}
