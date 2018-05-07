package controllers

import (
	"github.com/astaxie/beego"
	"storeManage/models"
	"storeManage/services"
)

type GoodsController struct {
	beego.Controller
}

// @Title Goods
// @Description 获取商品
// @Success 200 {object} []models.Goods "返回结果"
// @Failure 400 {object} []models.Goods "返回结果"
// @router /queryAll [POST]
func (m *MenuController) GetGoods() {
	var reply []models.Goods

	reply = services.GetGoods()
	m.Ctx.Output.Status = 200
	m.Data["json"] = reply
	m.ServeJSON()
}