package controllers

import (
	"github.com/astaxie/beego"
	"storeManage/models"
	"storeManage/services"
	"fmt"
	"storeManage/modelDB"
)

type OrderController struct {
	beego.Controller
}

// @Title Order
// @Description 获取订单
// @Param body body Page true "分页"
// @Success 200 {object} models.OrderResult "返回结果"
// @Failure 400 {object} models.OrderResult "返回结果"
// @router / [POST]
func (o *OrderController) GetOrder() {
	//var reply []models.Goods
	var reply models.OrderResult

	in := models.PageInfo{}
	o.ParseForm(&in)

	fmt.Println("PAGE:", in.Page)
	fmt.Println("SIZE:", in.Rows)

	reply = services.GetOrder(in)

	o.Ctx.Output.Status = 200
	o.Data["json"] = reply
	o.ServeJSON()
}

// @Title Order
// @Description 获取我的订单
// @Param body body Page true "分页"
// @Success 200 {object} models.OrderResult "返回结果"
// @Failure 400 {object} models.OrderResult "返回结果"
// @router /me [POST]
func (o *OrderController) GetMeOrder() {
	//var reply []models.Goods
	var reply models.OrderResult

	user := o.GetSession("user")
	use := user.(modelDB.User)

	in := models.PageInfo{}
	o.ParseForm(&in)

	fmt.Println("PAGE:", in.Page)
	fmt.Println("SIZE:", in.Rows)

	reply = services.GetMeOrder(in, use.Id)

	o.Ctx.Output.Status = 200
	o.Data["json"] = reply
	o.ServeJSON()
}

//生成订单
// @router /create [GET]
func (o *OrderController) CreateOrder() {
	var reply models.Result

	in := models.OrderRemarkParam{}
	o.ParseForm(&in)

	orderGoodsSession := o.GetSession("orderGoodsList")
	user := o.GetSession("user")
	use := user.(modelDB.User)
	if orderGoodsSession != nil {
		orderGoodsList := orderGoodsSession.(map[string]int)
		reply = services.CreateOrder(use.Id, orderGoodsList, in.Remark)
	}

	if reply.Success {
		o.DelSession("orderGoodsList")
	}

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

// 获取订单价格
// @router /getSum [GET]
func (o *OrderController) GetOrderSum() {
	//var reply []models.Goods
	var reply models.Result

	orderGoodsSession := o.GetSession("orderGoodsList")
	if orderGoodsSession != nil {
		orderGoodsList := orderGoodsSession.(map[string]int)
		reply = services.GetOrderSum(orderGoodsList)
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

// @Title OrderAddGoods
// @Description 订单删除商品
// @Param body body models.OrderAddGoodsParam true "订单添加商品"
// @Success 200 {object} []models.Result "返回结果"
// @Failure 400 {object} []models.Result "返回结果"
// @router /deleteGoods [POST]
func (o *OrderController) OrderDeleteGoods() {
	var reply models.Result

	in := models.IdParam{}
	o.ParseForm(&in)

	orderGoodsSession := o.GetSession("orderGoodsList")
	orderGoodsList := orderGoodsSession.(map[string]int)

	delete(orderGoodsList,in.Id)

	o.SetSession("orderGoodsList", orderGoodsList)

	o.Ctx.Output.Status = 200
	o.Data["json"] = reply
	o.ServeJSON()
}

