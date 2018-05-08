package controllers

import (
	"github.com/astaxie/beego"
	"storeManage/models"
	"storeManage/services"
	"fmt"
)

type OrderController struct {
	beego.Controller
}

// @Title Order
// @Description 获取订单
// @Param body body Page true "分页"
// @Success 200 {object} []models.Goods "返回结果"
// @Failure 400 {object} []models.Goods "返回结果"
// @router /queryAll [POST]
func (o *OrderController) GetOrder() {
	//var reply []models.Goods
	var reply models.GoodsResult

	in := models.PageInfo{}
	o.ParseForm(&in)

	fmt.Println("PAGE:", in.Page)
	fmt.Println("SIZE:", in.Rows)

	//reply.Goods = services.GetGoodss(in)
	reply.Total = 10
	o.Ctx.Output.Status = 200
	o.Data["json"] = reply
	o.ServeJSON()
}

// @Title NowOrder
// @Description 获取当前订单
// @Param body body Page true "分页"
// @Success 200 {object} models.OrderGoodsResult "返回结果"
// @Failure 400 {object} models.OrderGoodsResult "返回结果"
// @router /nowOrder [POST]
func (o *OrderController) GetNowOrder() {
	//var reply []models.Goods
	var reply models.OrderGoodsResult

	in := models.PageInfo{}
	o.ParseForm(&in)

	fmt.Println("PAGE:", in.Page)
	fmt.Println("SIZE:", in.Rows)

	orderGoodsSession := o.GetSession("orderGoodsList")
	if orderGoodsSession != nil {
		orderGoodsList := orderGoodsSession.(map[string]int)
		reply = services.GetNowOrder(in, orderGoodsList)
	}

	o.Ctx.Output.Status = 200
	o.Data["json"] = reply
	o.ServeJSON()
}

// @Title OrderAddGoods
// @Description 订单添加商品
// @Param body body models.OrderAddGoodsParam true "订单添加商品"
// @Success 200 {object} []models.Result "返回结果"
// @Failure 400 {object} []models.Result "返回结果"
// @router /addGoods [POST]
func (o *OrderController) OrderAddGoods() {
	var reply models.Result

	in := models.OrderAddGoodsParam{}
	o.ParseForm(&in)
	fmt.Println("id:",in.Id)
	fmt.Println("number:", in.Number)

	reply = services.CheckGoodsNumber(in.Id, in.Number)
	if reply.Success {
		orderGoodsList := make(map[string]int)
		orderGoodsSession := o.GetSession("orderGoodsList")
		if orderGoodsSession != nil {
			orderGoodsList = orderGoodsSession.(map[string]int)
		}
		orderGoodsList[in.Id] = in.Number
		o.SetSession("orderGoodsList", orderGoodsList)
	}

	o.Ctx.Output.Status = 200
	o.Data["json"] = reply
	o.ServeJSON()
}

