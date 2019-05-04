package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/jdc996/sudamericana_mid/models"
	"github.com/manucorporat/try"
	"github.com/udistrital/utils_oas/request"
)

// LugaresController operations for Ciudad
type LugaresController struct {
	beego.Controller
}

// URLMapping ...
func (c *LugaresController) URLMapping() {
	c.Mapping("GetLugares", c.GetLugares)
}

// GetLugares ...
// @Title Get Lugares
// @Description get Lugares
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 {object} models.Ciudad
// @Failure 403 not found resource
// @router getLugares/ [get]
func (c *LugaresController) GetLugares() {

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
	c.ServeJSON()
	/*resp, error := http.Get("localhost:8080/ciudad")

	if error != nil {
		logs.Error(error)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = error
		c.Abort("404")
	} else {
		c.Data["json"] = resp
		c.ServeJSON()
	}
	*/
	// if err := request.GetJsonWSO2(beego.AppConfig.String("Wso2Service")+"servicios_academicos/consulta_datos_docente_planta/"+Id, &resLugar); err == nil {
	// 	if resLugar["datosCollection"].(map[string]interface{})["datos"] != nil {
	// 		c.Data["json"] = resLugar["datosCollection"].(map[string]interface{})["datos"].([]interface{})[0]
	// 	}
	// } else {
	// 	//c.Data["json"] = models.Alert{Code: "E_0458", Body: err.Error(), Type: "error"}
	// 	//beego.Error("Error", err.Error())
	// 	return
	// }
}

//PostPais ...
// @Title Post Pais
// @Descripcion post pais
// @Param	body		body 	models.Pais	true	"body for Postpais content"
// @Success 200 {object} models.Pais
// @Failure 403 body is empty
// @router /postPais [post]
func (c *LugaresController) PostPais() {
	beego.Info("entro a post pais")
	try.This(func() {
		var v models.Pais
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			Urlcrud := beego.AppConfig.String("UrlLugaresCrud") + "pais"
			beego.Info("URL ", Urlcrud)
			if err := request.SendJson(Urlcrud, "POST", nil, &v); err == nil {
				c.Data["json"] = err
			} else {
				panic(err.Error())
			}
		} else {
			c.Data["json"] = err.Error()
		}
	}).Catch(func(e try.E) {
		beego.Info("excep: ", e)
		//c.Data["json"] = models.Alert{Code: "E_0458", Body: e, Type: "error"}
	})
	c.ServeJSON()
}

/*
func (c *LugaresController) GetAll() {
	idStr := c.Ctx.Input.Param(":id")

	resp, error := http.Get("http://localhost:8080/ciudad/id=" + idStr)

	if error != nil {
		logs.Error(error)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = error
		c.Abort("404")
	} else {
		c.Data["json"] = resp
		c.ServeJSON()
	}

}
*/
