package controllers

import (
	"github.com/astaxie/beego"
	"storeManage/models"
	"storeManage/services"
	"fmt"
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

// @Title TypePage
// @Description 获取分类（分页）
// @Param body body Page true "分页"
// @Success 200 {object} models.GoodsResult "返回结果"
// @Failure 400 {object} models.GoodsResult "返回结果"
// @router /queryAll [POST]
func (g *TypeController) GetTypePage() {
	//var reply []models.Goods
	var reply models.TypePageResult

	in := models.PageInfo{}
	g.ParseForm(&in)

	fmt.Println("PAGE:", in.Page)
	fmt.Println("SIZE:", in.Rows)

	reply = services.GetTypes(in)
	g.Ctx.Output.Status = 200
	g.Data["json"] = reply
	g.ServeJSON()
}

// @Title SaveTypeEditId
// @Description 获取修改商品ID
// @Param body body models.SaveEditId true "修改商品ID"
// @Success 200 {object} []models.Result "返回结果"
// @Failure 400 {object} []models.Result "返回结果"
// @router /saveEditId [POST]
func (g *TypeController) SaveTypeEditId() {
	var reply models.Result

	in := models.IdParam{}
	g.ParseForm(&in)
	fmt.Println("SaveTypeEditId:",in.Id)
	g.SetSession("typeEditId", in.Id)

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
func (g *TypeController) GetEditType() {
	var reply models.Type

	typeId := g.GetSession("typeEditId")
	fmt.Print("goodsId:", typeId)

	reply = services.GetTypeById(typeId.(string))

	g.Ctx.Output.Status = 200
	g.Data["json"] = reply
	g.ServeJSON()
}

// @Title EditType
// @Description 修改商品信息
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /edit [POST]
func (g *TypeController) EditType() {
	var reply models.Result

	typeId := g.GetSession("typeEditId")
	fmt.Print("typeEditId:", typeId)

	in := models.GoodsParam{}
	g.ParseForm(&in)

	goodsInfo :=  make(map[string]interface{})
	goodsInfo["NAME"] = in.Name
	goodsInfo["PARENT_ID"] = in.TypeId

	reply = services.EditType(typeId.(string), goodsInfo)

	g.Ctx.Output.Status = 200
	g.Data["json"] = reply
	g.ServeJSON()
}

// @Title SaveGoods
// @Description 添加商品
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /save [POST]
func (g *TypeController) SaveType() {
	var reply models.Result

	in := models.GoodsParam{}
	g.ParseForm(&in)

	reply = services.SaveType(in)

	g.Ctx.Output.Status = 200
	g.Data["json"] = reply
	g.ServeJSON()
}

// @Title DeleteGoods
// @Description 删除商品
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /remove [POST]
func (g *TypeController) DeleteType() {
	var reply models.Result

	in := models.IdParam{}
	g.ParseForm(&in)

	reply = services.DeleteType(in.Id)

	g.Ctx.Output.Status = 200
	reply.Success = true
	reply.Message = "删除商品成功:ID=" + in.Id
	g.Data["json"] = reply
	g.ServeJSON()
}