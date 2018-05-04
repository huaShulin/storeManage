package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"storeManage/models"
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
func (l *LoginController) UserLoginVerification() {
	var reply models.Result

	fmt.Println("UserLoginVerification")
	login := models.Login{}
	if err := l.ParseForm(&login); err != nil {
		//l.Ctx.WriteString("输入参数不合法")
		l.Ctx.Output.Status = 200
		reply.Success = false
		reply.Message = "输入参数不合法"
		l.Data["json"] = reply
		l.ServeJSON()
		return
	}
	if login.Username != "shulin" {
		//l.Ctx.WriteString("用户不存在")
		l.Ctx.Output.Status = 200
		reply.Success = false
		reply.Message = "用户不存在"
		l.Data["json"] = reply
		l.ServeJSON()
		return
	}
	if login.Password != "123123" {
		//l.Ctx.WriteString("密码错误")
		l.Ctx.Output.Status = 200
		reply.Success = false
		reply.Message = "密码错误"
		l.Data["json"] = reply
		l.ServeJSON()
		return
	}
	//l.Ctx.WriteString("登陆成功")

	l.Ctx.Output.Status = 200

	l.Data["json"] = reply
	l.ServeJSON()

}