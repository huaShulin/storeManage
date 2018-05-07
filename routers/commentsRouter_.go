package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["storeManage/controllers:LoginController"] = append(beego.GlobalControllerRouter["storeManage/controllers:LoginController"],
		beego.ControllerComments{
			Method: "GetHelloWorld",
			Router: `/GetHelloWorld`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["storeManage/controllers:LoginController"] = append(beego.GlobalControllerRouter["storeManage/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/Login`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["storeManage/controllers:MenuController"] = append(beego.GlobalControllerRouter["storeManage/controllers:MenuController"],
		beego.ControllerComments{
			Method: "GetGoods",
			Router: `/queryAll`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["storeManage/controllers:MenuController"] = append(beego.GlobalControllerRouter["storeManage/controllers:MenuController"],
		beego.ControllerComments{
			Method: "GetMenu",
			Router: `/`,
			AllowHTTPMethods: []string{"GET"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["storeManage/controllers:MenuController"] = append(beego.GlobalControllerRouter["storeManage/controllers:MenuController"],
		beego.ControllerComments{
			Method: "GetGoods",
			Router: `/`,
			AllowHTTPMethods: []string{"GET"},
			MethodParams: param.Make(),
			Params: nil})

}
