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
// @Success 200 {object} models.GoodsResult "返回结果"
// @Failure 400 {object} models.GoodsResult "返回结果"
// @router /queryAll [POST]
func (g *GoodsController) GetGoods() {
	//var reply []models.Goods
	var reply models.GoodsResult

	in := models.PageInfo{}
	g.ParseForm(&in)

	fmt.Println("PAGE:", in.Page)
	fmt.Println("SIZE:", in.Rows)

	reply = services.GetGoodss(in)
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

	in := models.IdParam{}
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

// @Title EditGoods
// @Description 修改商品信息
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /edit [POST]
func (g *GoodsController) EditGoods() {
	var reply models.Result

	goodsId := g.GetSession("goodsEditId")
	fmt.Print("goodsId:", goodsId)

	in := models.GoodsParam{}
	g.ParseForm(&in)

	goodsInfo :=  make(map[string]interface{})
	goodsInfo["NAME"] = in.Name
	goodsInfo["PRICE"] = in.Price
	goodsInfo["NUMBER"] = in.Number
	goodsInfo["PURCHASE"] = in.Purchase
	goodsInfo["REMARK"] = in.Remark
	goodsInfo["TYPE_ID"] = in.TypeId

	reply = services.EditGoods(goodsId.(string), goodsInfo)

	g.Ctx.Output.Status = 200
	g.Data["json"] = reply
	g.ServeJSON()
}

// @Title SaveGoods
// @Description 添加商品
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /save [POST]
func (g *GoodsController) SaveGoods() {
	var reply models.Result

	in := models.GoodsParam{}
	g.ParseForm(&in)

	reply = services.SaveGoods(in)

	g.Ctx.Output.Status = 200
	g.Data["json"] = reply
	g.ServeJSON()
}

// @Title DeleteGoods
// @Description 删除商品
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /remove [POST]
func (g *GoodsController) DeleteGoods() {
	var reply models.Result

	in := models.IdParam{}
	g.ParseForm(&in)

	reply = services.DeleteGoods(in.Id)

	g.Ctx.Output.Status = 200
	reply.Success = true
	reply.Message = "删除商品成功:ID=" + in.Id
	g.Data["json"] = reply
	g.ServeJSON()
}