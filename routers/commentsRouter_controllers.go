package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/jdc996/sudamericana_mid/controllers:EquipoController"] = append(beego.GlobalControllerRouter["github.com/jdc996/sudamericana_mid/controllers:EquipoController"],
        beego.ControllerComments{
            Method: "RegistrarEquipo",
            Router: `/registrarEquipo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jdc996/sudamericana_mid/controllers:EquipoController"] = append(beego.GlobalControllerRouter["github.com/jdc996/sudamericana_mid/controllers:EquipoController"],
        beego.ControllerComments{
            Method: "GetEquipos",
            Router: `getEquipos/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jdc996/sudamericana_mid/controllers:EquipoController"] = append(beego.GlobalControllerRouter["github.com/jdc996/sudamericana_mid/controllers:EquipoController"],
        beego.ControllerComments{
            Method: "GetLugar",
            Router: `getLugar/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jdc996/sudamericana_mid/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/jdc996/sudamericana_mid/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "PostPais",
            Router: `/postPais`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jdc996/sudamericana_mid/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/jdc996/sudamericana_mid/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "GetLugares",
            Router: `getLugares/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
