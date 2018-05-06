package controllers

import (
	"github.com/astaxie/beego"
	"storeManage/models"
	"storeManage/services"
	"fmt"
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

	userId := m.GetSession("id")
	fmt.Print("获取菜单",userId)

	s := userId.(string)
	fmt.Print("获取菜单",s)

	reply = services.Menu(s)
	m.Ctx.Output.Status = 200
	m.Data["json"] = reply
	m.ServeJSON()
}