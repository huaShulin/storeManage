package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["storeManage/controllers:GoodsController"] = append(beego.GlobalControllerRouter["storeManage/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "GetGoods",
			Router: `/queryAll`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["storeManage/controllers:GoodsController"] = append(beego.GlobalControllerRouter["storeManage/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "SaveEditId",
			Router: `/saveEditId`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["storeManage/controllers:GoodsController"] = append(beego.GlobalControllerRouter["storeManage/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "GetEditGoods",
			Router: `/getEditGoods`,
			AllowHTTPMethods: []string{"GET"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["storeManage/controllers:GoodsController"] = append(beego.GlobalControllerRouter["storeManage/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "EditGoods",
			Router: `/edit`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["storeManage/controllers:GoodsController"] = append(beego.GlobalControllerRouter["storeManage/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "SaveGoods",
			Router: `/save`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["storeManage/controllers:GoodsController"] = append(beego.GlobalControllerRouter["storeManage/controllers:GoodsController"],
		beego.ControllerComments{
			Method: "DeleteGoods",
			Router: `/remove`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams: param.Make(),
			Params: nil})

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
			Method: "GetMenu",
			Router: `/`,
			AllowHTTPMethods: []string{"GET"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["storeManage/controllers:OrderController"] = append(beego.GlobalControllerRouter["storeManage/controllers:OrderController"],
		beego.ControllerComments{
			Method: "GetOrder",
			Router: `/queryAll`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["storeManage/controllers:OrderController"] = append(beego.GlobalControllerRouter["storeManage/controllers:OrderController"],
		beego.ControllerComments{
			Method: "GetNowOrder",
			Router: `/nowOrder`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["storeManage/controllers:OrderController"] = append(beego.GlobalControllerRouter["storeManage/controllers:OrderController"],
		beego.ControllerComments{
			Method: "OrderAddGoods",
			Router: `/addGoods`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["storeManage/controllers:TypeController"] = append(beego.GlobalControllerRouter["storeManage/controllers:TypeController"],
		beego.ControllerComments{
			Method: "GetType",
			Router: `/`,
			AllowHTTPMethods: []string{"GET"},
			MethodParams: param.Make(),
			Params: nil})

}
