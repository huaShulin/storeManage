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
			Method: "UserLoginVerification",
			Router: `/Login`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams: param.Make(),
			Params: nil})

}
