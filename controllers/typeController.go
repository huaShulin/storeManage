package controllers

import (
	"github.com/astaxie/beego"
	"storeManage/models"
	"storeManage/services"
)

type TypeController struct {
	beego.Controller
}

// @Title Type
// @Description 获取分类
// @Success 200 {object} []models.Type "返回结果"
// @Failure 400 {object} []models.Type "返回结果"
// @router / [GET]
func (t *TypeController) GetType() {
	var reply []models.Type

	reply = services.Type()
	t.Ctx.Output.Status = 200
	t.Data["json"] = reply
	t.ServeJSON()
}