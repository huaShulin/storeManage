package controllers

import (
	"github.com/astaxie/beego"
	"storeManage/models"
	"storeManage/services"
	"fmt"
	"storeManage/modelDB"
)

type MenuController struct {
	beego.Controller
}

// @Title Menu
// @Description 获取菜单
// @Success 200 {object} models.MenuResult "返回结果"
// @Failure 400 {object} models.MenuResult "返回结果"
// @router / [GET]
func (m *MenuController) GetMenu() {
	var reply models.MenuResult

	user := m.GetSession("user")

	if user == nil {
		m.Ctx.Output.Status = 200
		reply.Success = false
		reply.Message = "用户未登录"
		m.Data["json"] = reply
		m.ServeJSON()
		return
	}

	use := user.(modelDB.User)
	fmt.Print("获取菜单",use)

	reply = services.MenuByUser(use.Id)
	m.Ctx.Output.Status = 200
	m.Data["json"] = reply
	m.ServeJSON()
}

// @Title Menu
// @Description 获取菜单
// @Success 200 {object} models.MenuResult "返回结果"
// @Failure 400 {object} models.MenuResult "返回结果"
// @router /child [GET]
func (m *MenuController) GetChildMenu() {
	var reply []models.Menu

	reply = services.MenuChild()
	fmt.Println("reply:",reply)
	m.Ctx.Output.Status = 200
	m.Data["json"] = reply
	m.ServeJSON()
}