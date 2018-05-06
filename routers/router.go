package routers

import (
	"storeManage/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	ns :=
		beego.NewNamespace("/hyl",
			beego.NSNamespace("/login",
				beego.NSInclude(
					&controllers.LoginController{},
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
		)
	beego.AddNamespace(ns)
}
