package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/jdc996/sudamericana_mid/models"
	"github.com/manucorporat/try"
	"github.com/udistrital/utils_oas/request"
)

// PartidosController operations for partidos
type PartidosController struct {
	beego.Controller
}

// URLMapping ...
func (c *PartidosController) URLMapping() {
	c.Mapping("GetPartidos", c.GetPartidos)
	c.Mapping("RegistrarEquipos", c.RegistrarPartido)
}

// GetPartidos ...
// @Title Get Partidos
// @Description get Partidos
// @Success 200 {object} models.Partidos
// @Failure 403 not found resource
// @router GetPartidos/ [get]
func (c *PartidosController) GetPartidos() {
	//var queryStr string
	beego.Info("entra")
	//if r := c.GetString("query"); r != "" {
	//	queryStr = r
	//}
	beego.Info(beego.AppConfig.String("UrlPartidosCrud") + "partidos/")
	var partidos []models.Partidos
	//queryFilter := "?query=id%3A"

	if err := request.GetJson(beego.AppConfig.String("UrlPartidosCrud")+"partidos", &partidos); err == nil {
		beego.Info("retorna")
		c.Data["json"] = partidos

	} else {
		c.Data["system"] = err
		c.Abort("404")
	}
	beego.Info()
	c.ServeJSON()
}

//RegistrarPartido ...
// @Title Registrar Partido
// @Descripcion registrar Partido
// @Param	body		body 	models.Partido	true	"body for registrar partido content"
// @Success 200 {object} models.Partidos
// @Failure 403 body is empty
// @router /registrarPartido [post]
func (c *PartidosController) RegistrarPartido() {
	try.This(func() {
		var body models.Partidos
		var res map[string]interface{}
		var idEquipoLocal, idEquipoVisitante string
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err == nil {
			idEquipoLocal = fmt.Sprintf("%v", body.EquipoLocalId)
			idEquipoVisitante = fmt.Sprintf("%v", body.EquipoVisitanteId)
			beego.Info("idEquipo local ", idEquipoLocal)
			beego.Info("idEquipo visitante", idEquipoVisitante)
			beego.Info("entro a Registrar Partido")
			equipoLocal := c.GetEquipoById(idEquipoLocal)[0]
			equipoVisitante := c.GetEquipoById(idEquipoVisitante)[0]
			beego.Info("Equipo local id: ", equipoLocal.Id)
			beego.Info("Equipo visitante id: ", equipoVisitante.Id)
			if equipoLocal.Id != 0 && equipoVisitante.Id != 0 {
				beego.Info("existen los equipos")
				Urlcrud := beego.AppConfig.String("UrlPartidosCrud") + "partidos"
				beego.Info("URL ", Urlcrud)
				posicionLocal := c.GetPocisionByEquipoId(strconv.Itoa(equipoLocal.Id))[0]
				beego.Info("posicion equipo local: ", posicionLocal)
				posicionVisitante := c.GetPocisionByEquipoId(strconv.Itoa(equipoVisitante.Id))[0]
				beego.Info("posicion equipo visitante: ", posicionVisitante)
				posicionLocal.GolesFavor += body.GolesLocal
				posicionLocal.GolesContra += body.GolesVisitante
				posicionLocal.PartidosJugados++
				posicionVisitante.GolesFavor += body.GolesVisitante
				posicionVisitante.GolesContra += body.GolesLocal
				posicionVisitante.PartidosJugados++
				posicionLocal.DiferenciaGoles = posicionLocal.GolesFavor - posicionLocal.GolesContra
				posicionVisitante.DiferenciaGoles = posicionVisitante.GolesFavor - posicionVisitante.GolesContra
				if body.GolesLocal > body.GolesVisitante {
					posicionLocal.PartidosGanados++
					posicionLocal.Puntaje += 3
					posicionVisitante.PartidosPerdidos++

				} else if body.GolesVisitante > body.GolesLocal {

					posicionVisitante.PartidosGanados++
					posicionLocal.PartidosPerdidos++
					posicionVisitante.Puntaje += 3
				} else {
					posicionLocal.PartidosEmpatados++
					posicionVisitante.PartidosEmpatados++
					posicionLocal.Puntaje++
					posicionVisitante.Puntaje++
				}
				beego.Info("posicion equipo local tras partido: ", posicionLocal)
				beego.Info("posicion equipp Visitante tras partido:", posicionVisitante)
				if err := request.SendJson(Urlcrud, "POST", &res, &body); err == nil {

					beego.Info("Post data: ", res)
					c.Data["json"] = res
					// 				for data := range res {
					// 					beego.Info("data res:", data)
					// 					beego.Info(res[data])
					// 				}

					// 	// 				beego.Info("posicion: ", posicion)
					c.ActualizarPosicion(posicionLocal)
					c.ActualizarPosicion(posicionVisitante)
				} else {
					panic(err.Error())
				}
			} else {
				beego.Info("no existen lo(s) equipo(s)")
				c.Data["json"] = models.AlertString{Type: "error", Code: "E_0458", Body: "No existen lo(s) equipo(s)"}
			}
		} else {
			c.Data["json"] = err.Error()
		}
	}).Catch(func(e try.E) {
		beego.Info("excep: ", e)
		c.Data["json"] = models.AlertError{Code: "E_0458", Body: e, Type: "error"}
	})
	c.ServeJSON()
}

func (c *PartidosController) ActualizarPosicion(posicion models.Posicion) (res map[string]interface{}) {
	beego.Info("entro a Actualizar posiciones")

	pos := posicion
	//values := map[string]int{"Id": 0, "PartidosJugados": 0, "EquipoId": idEquipoiInt, "PartidosEmpatados": 0, "PartidosGanados": 0, "PartidosPerdidos": 0, "GolesFavor": 0, "GolesContra": 0, "DiferenciaGoles": 0}
	//jsonValue, _ := json.Marshal(values)
	Urlcrud := beego.AppConfig.String("UrlLigaCrud") + "posicion/" + strconv.Itoa(pos.Id)
	beego.Info(Urlcrud)

	if err := request.SendJson(Urlcrud, "PUT", &res, &pos); err == nil {
		beego.Info("Actualizar posiciones: ", res)
		beego.Info("Retorna")
	} else {
		beego.Info("error: ", err)
	}

	beego.Info("Res Actualizar info:", res)
	return
}

func (c *PartidosController) GetPocisionByEquipoId(idEquipo string) (res []models.Posicion) {
	if err := request.GetJson(beego.AppConfig.String("UrlLigaCrud")+"posicion/?query=EquipoId:"+idEquipo, &res); err == nil {
		beego.Info("Retorna posicion")
	} else {
		beego.Info("error")
	}
	beego.Info("res", res)
	return
}

func (c *PartidosController) GetEquipoById(idEquipo string) (res []models.Equipo) {
	if err := request.GetJson(beego.AppConfig.String("UrlEquipoCrud")+"equipo/?query=Id:"+idEquipo, &res); err == nil {
		beego.Info("Retorna")
	} else {
		beego.Info("error")
	}
	beego.Info("res", res)
	return
}
