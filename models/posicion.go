package models

type Posicion struct {
	Id                int
	Puntaje           int
	EquipoId          int
	PartidosJugados   int
	PartidosGanados   int
	PartidosEmpatados int
	PartidosPerdidos  int
	GolesFavor        int
	GolesContra       int
	DiferenciaGoles   int
}
