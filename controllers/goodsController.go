package controllers

import (
	"github.com/astaxie/beego"
	"storeManage/models"
	"storeManage/services"
	"fmt"
)

type GoodsController struct {
	beego.Controller
}

// @Title Goods
// @Description 获取商品
// @Param body body Page true "分页"
// @Success 200 {object} []models.Goods "返回结果"
// @Failure 400 {object} []models.Goods "返回结果"
// @router /queryAll [POST]
func (g *GoodsController) GetGoods() {
	//var reply []models.Goods
	var reply models.GoodsResult

	in := models.PageInfo{}
	g.ParseForm(&in)

	fmt.Println("PAGE:", in.Page)
	fmt.Println("SIZE:", in.Rows)

	reply.Goods = services.GetGoodss()
	reply.Total = 10
	g.Ctx.Output.Status = 200
	g.Data["json"] = reply
	g.ServeJSON()
}

// @Title SaveEditId
// @Description 获取修改商品ID
// @Param body body models.SaveEditId true "修改商品ID"
// @Success 200 {object} []models.Result "返回结果"
// @Failure 400 {object} []models.Result "返回结果"
// @router /saveEditId [POST]
func (g *GoodsController) SaveEditId() {
	var reply models.Result

	in := models.SaveEditId{}
	g.ParseForm(&in)
	fmt.Println("SaveEditId:",in.Id)
	g.SetSession("goodsEditId", in.Id)

	g.Ctx.Output.Status = 200
	reply.Success = true
	reply.Message = "目标商品存储成功:ID=" + in.Id
	g.Data["json"] = reply
	g.ServeJSON()
}

// @Title GetEditGoods
// @Description 获取修改商品
// @Success 200 {object} []models.Result "返回结果"
// @Failure 400 {object} []models.Result "返回结果"
// @router /getEditGoods [GET]
func (g *GoodsController) GetEditGoods() {
	var reply models.Goods

	goodsId := g.GetSession("goodsEditId")
	fmt.Print("goodsId:", goodsId)

	reply = services.GetGoods(goodsId.(string))

	g.Ctx.Output.Status = 200
	g.Data["json"] = reply
	g.ServeJSON()
}