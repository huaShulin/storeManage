package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"storeManage/models"
	"storeManage/services"
	"storeManage/modelDB"
)

type UserController struct {
	beego.Controller
}

// @router /GetHelloWorld
func (u *UserController) GetHelloWorld() {
	fmt.Println("Hello World")
	u.Data["Website"] = "Hello"
	u.Data["Email"] = "World"
	//h.TplName = "index.tpl"
	u.Data["json"] = "Hello World"
	u.ServeJSON()
}

// @Title Login
// @Description 登录
// @Param body body models.Login true "登录结构体"
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /login [POST]
func (u *UserController) Login() {
	var reply models.Result

	fmt.Println("Login")
	login := models.Login{}
	if err := u.ParseForm(&login); err != nil {
		u.Ctx.Output.Status = 200
		reply.Success = false
		reply.Message = "输入参数不合法"
		u.Data["json"] = reply
		u.ServeJSON()
		return
	}

	var user modelDB.User
	reply, user = services.Login(login)

	if reply.Success {
		u.SetSession("user", user)
		fmt.Println(u.GetSession("id"))
	}

	u.Ctx.Output.Status = 200
	u.Data["json"] = reply
	u.ServeJSON()
}

// @Title User
// @Description 获取用户
// @Param body body Page true "分页"
// @Success 200 {object} models.GoodsResult "返回结果"
// @Failure 400 {object} models.GoodsResult "返回结果"
// @router /queryAll [POST]
func (u *UserController) GetUser() {
	//var reply []models.Goods
	var reply models.UserResult

	in := models.PageInfo{}
	u.ParseForm(&in)

	fmt.Println("PAGE:", in.Page)
	fmt.Println("SIZE:", in.Rows)

	reply = services.GetUsers(in)
	u.Ctx.Output.Status = 200
	u.Data["json"] = reply
	u.ServeJSON()
}

// @Title StatusUser
// @Description 修改用户状态
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /status [POST]
func (u *UserController) StatusUser() {
	var reply models.Result

	in := models.IdParam{}
	u.ParseForm(&in)

	reply = services.StatusUser(in.Id)

	u.Ctx.Output.Status = 200
	reply.Success = true
	reply.Message = "修改用户状态成功:ID=" + in.Id
	u.Data["json"] = reply
	u.ServeJSON()
}

// @Title DeleteUser
// @Description 删除用户
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /remove [POST]
func (u *UserController) DeleteUser() {
	var reply models.Result

	in := models.IdParam{}
	u.ParseForm(&in)

	reply = services.DeleteUser(in.Id)

	u.Ctx.Output.Status = 200
	reply.Success = true
	reply.Message = "删除用户成功:ID=" + in.Id
	u.Data["json"] = reply
	u.ServeJSON()
}