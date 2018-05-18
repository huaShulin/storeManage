package routers

import (
	"storeManage/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	ns :=
		beego.NewNamespace("/hyl",
			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
				),
			),
			beego.NSNamespace("/menu",
				beego.NSInclude(
					&controllers.MenuController{},
				),
			),
			beego.NSNamespace("/goods",
				beego.NSInclude(
					&controllers.GoodsController{},
				),
			),
			beego.NSNamespace("/type",
				beego.NSInclude(
					&controllers.TypeController{},
				),
			),
			beego.NSNamespace("/order",
				beego.NSInclude(
					&controllers.OrderController{},
				),
			),
		)
	beego.AddNamespace(ns)
}
