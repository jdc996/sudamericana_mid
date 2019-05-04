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

// LugaresController operations for equipo
type EquipoController struct {
	beego.Controller
}

// URLMapping ...
func (c *EquipoController) URLMapping() {
	c.Mapping("GetEquipos", c.GetEquipos)
	c.Mapping("RegistrarEquipo", c.RegistrarEquipo)
}

// GetEquipos ...
// @Title Get Equipos
// @Description get Equipos
// @Success 200 {object} models.Equipo
// @Failure 403 not found resource
// @router getEquipos/ [get]
func (c *EquipoController) GetEquipos() {
	//idStr := c.Ctx.Input.Param(":id")
	beego.Info("entra")
	beego.Info(beego.AppConfig.String("UrlEquipoCrud") + "equipo")
	var equipos []models.Equipo
	if err := request.GetJson(beego.AppConfig.String("UrlEquipoCrud")+"equipo", &equipos); err == nil {
		beego.Info("retorna")
		c.Data["json"] = equipos
	} else {
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

//RegistrarEquipo ...
// @Title Registrar equipo
// @Descripcion registrar equipo
// @Param	body		body 	models.Equipo	true	"body for Postpais content"
// @Success 200 {object} models.Equipo
// @Failure 403 body is empty
// @router /registrarEquipo [post]
func (c *EquipoController) RegistrarEquipo() {
	try.This(func() {
		var body models.Equipo
		var res map[string]interface{}
		var idCiudad string
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err == nil {
			idCiudad = fmt.Sprintf("%v", body.CiudadId)
			beego.Info("idCIudad: ", idCiudad)
			beego.Info("entro a Registrar Equipo")
			ciudad := c.GetCiudadById(idCiudad)[0]
			beego.Info("ciudades", ciudad)
			if ciudad.Id != 0 {
				beego.Info("existe la ciudad")
				Urlcrud := beego.AppConfig.String("UrlEquipoCrud") + "equipo"
				beego.Info("URL ", Urlcrud)

				if err := request.SendJson(Urlcrud, "POST", &res, &body); err == nil {

					beego.Info("Post data: ", res)
					c.Data["json"] = res
					for data := range res {
						beego.Info("data res:", data)
						beego.Info(res[data])
					}
					idEquipo := fmt.Sprintf("%v", res["Id"])
					posicion := c.PostPosiciones(idEquipo)
					beego.Info("posicion: ", posicion)

				} else {
					panic(err.Error())
				}
			} else {
				beego.Info("no ciudad existe")
				c.Data["json"] = models.AlertString{Type: "error", Code: "E_0458", Body: "No existe la ciudad"}
			}
		} else {
			c.Data["json"] = err.Error()
		}
	}).Catch(func(e try.E) {
		beego.Info("excep: ", e)
		c.Data["json"] = models.AlertError{Code: "E_0458", Body: e, Type: "error"}
	})
	c.ServeJSON()
	// beego.Info("id ciudad: ", idCiudad)
	// beego.Info("entro a Registrar Equipo")
	// ciudad := c.GetCiudadById(idCiudad)[0]
	// beego.Info("ciudades", ciudad)
	// if ciudad.Id != 0 {
	// 	beego.Info("existe la ciudad")
	// 	try.This(func() {
	// 		var equipo models.Equipo
	// 		Urlcrud := beego.AppConfig.String("UrlEquipoCrud") + "equipo"
	// 		beego.Info("URL ", Urlcrud)
	// 		request.SendJson(Urlcrud, "POST", nil, &equipo)
	// 		c.ServeJSON()
	// 	}).Catch(func(e try.E) {
	// 		beego.Info("excep: ", e)
	// 		c.Data["json"] = models.AlertError{Code: "E_0458", Body: e, Type: "error"}
	// 	})
	// } else {
	// 	beego.Info("no ciudad existe")
	// 	c.Data["json"] = models.AlertString{Type: "error", Code: "E_0458", Body: "No existe la ciudad"}
	// }

	// try.This(func() {
	// 	var v models.Pais
	// 	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
	// 		Urlcrud := beego.AppConfig.String("UrlLugaresCrud") + "pais"
	// 		beego.Info("URL ", Urlcrud)
	// 		if err := request.SendJson(Urlcrud, "POST", nil, &v); err == nil {
	// 			c.Data["json"] = err
	// 		} else {
	// 			panic(err.Error())
	// 		}
	// 	} else {
	// 		c.Data["json"] = err.Error()
	// 	}
	// }).Catch(func(e try.E) {
	// 	beego.Info("excep: ", e)
	// 	//c.Data["json"] = models.Alert{Code: "E_0458", Body: e, Type: "error"}
	// })

}

func (c *EquipoController) GetCiudadById(idCiudad string) (res []models.Ciudad) {
	if err := request.GetJson(beego.AppConfig.String("UrlLugaresCrud")+"ciudad/?query=Id:"+idCiudad, &res); err == nil {
		beego.Info("Retorna")
	} else {
		beego.Info("error")
	}
	beego.Info("res", res)
	return
}

func (c *EquipoController) PostPosiciones(idEquipo string) (res map[string]interface{}) {
	beego.Info("entro a Post posiciones")
	var posicion models.Posicion
	idEquipoiInt, nil := strconv.Atoi(idEquipo)
	posicion.Id = 0
	posicion.Puesto = 1

	posicion.EquipoId = idEquipoiInt
	posicion.PartidosJugados = 0
	posicion.PartidosEmpatados = 0
	posicion.PartidosGanados = 0
	posicion.PartidosPerdidos = 0
	posicion.GolesFavor = 0
	posicion.GolesContra = 0
	posicion.DiferenciaGoles = 0

	//values := map[string]int{"Id": 0, "PartidosJugados": 0, "EquipoId": idEquipoiInt, "PartidosEmpatados": 0, "PartidosGanados": 0, "PartidosPerdidos": 0, "GolesFavor": 0, "GolesContra": 0, "DiferenciaGoles": 0}
	//jsonValue, _ := json.Marshal(values)
	Urlcrud := beego.AppConfig.String("UrlLigaCrud") + "posicion"
	beego.Info(Urlcrud)
	if err := request.SendJson(Urlcrud, "POST", &res, &posicion); err == nil {
		beego.Info("Post data: ", res)
		beego.Info("Retorna")
	} else {
		beego.Info("error: ", err)
	}

	beego.Info("res", res)
	return
}

// GetLugar ...
// @Title Get Lugar
// @Description get Lugar
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 res models.Ciudad
// @Failure 403 not found resource
// @router getLugar/ [get]
func (c *EquipoController) GetLugar() (res []models.Ciudad) {
	var queryStr string
	beego.Info("entra")
	if r := c.GetString("query"); r != "" {
		queryStr = r
	}
	beego.Info(beego.AppConfig.String("UrlLugaresCrud") + "ciudad/" + queryStr)
	var ciudades []models.Ciudad
	//queryFilter := "?query=id%3A"
	beego.Info("queryStr= " + queryStr)
	if err := request.GetJson(beego.AppConfig.String("UrlLugaresCrud")+"ciudad/?query="+queryStr, &ciudades); err == nil {
		beego.Info("retorna")
		c.Data["json"] = ciudades

	} else {
		c.Data["system"] = err
		c.Abort("404")
	}
	beego.Info()
	//c.ServeJSON()
	return ciudades
}
