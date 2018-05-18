package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	//c.TplName = "/static/back/manager/managerLogin.html"
	c.Redirect("/static/back/user/userLogin.html",301)
}
