package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"storeManage/models"
	"storeManage/services"
)

type LoginController struct {
	beego.Controller
}

// @router /GetHelloWorld
func (l *LoginController) GetHelloWorld() {
	fmt.Println("Hello World")
	l.Data["Website"] = "Hello"
	l.Data["Email"] = "World"
	//h.TplName = "index.tpl"
	l.Data["json"] = "Hello World"
	l.ServeJSON()
}

// @Title Login
// @Description 登录
// @Param body body models.Login true "登录结构体"
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /Login [POST]
func (l *LoginController) Login() {
	var reply models.Result

	fmt.Println("Login")
	login := models.Login{}
	if err := l.ParseForm(&login); err != nil {
		l.Ctx.Output.Status = 200
		reply.Success = false
		reply.Message = "输入参数不合法"
		l.Data["json"] = reply
		l.ServeJSON()
		return
	}

	var id string
	reply, id = services.Login(login)

	if reply.Success {
		l.SetSession("id", id)
		fmt.Println(l.GetSession("id"))
	}

	l.Ctx.Output.Status = 200
	l.Data["json"] = reply
	l.ServeJSON()

}