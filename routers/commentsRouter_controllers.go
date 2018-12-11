package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-bee-restful/controllers:CarController"] = append(beego.GlobalControllerRouter["go-bee-restful/controllers:CarController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-bee-restful/controllers:CarController"] = append(beego.GlobalControllerRouter["go-bee-restful/controllers:CarController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-bee-restful/controllers:CarController"] = append(beego.GlobalControllerRouter["go-bee-restful/controllers:CarController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:Carid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-bee-restful/controllers:CarController"] = append(beego.GlobalControllerRouter["go-bee-restful/controllers:CarController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:Carid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-bee-restful/controllers:CarController"] = append(beego.GlobalControllerRouter["go-bee-restful/controllers:CarController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:Carid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
