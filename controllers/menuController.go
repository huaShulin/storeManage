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
	use := user.(modelDB.User)
	fmt.Print("获取菜单",use)

	reply = services.Menu(use.Id)
	m.Ctx.Output.Status = 200
	m.Data["json"] = reply
	m.ServeJSON()
}