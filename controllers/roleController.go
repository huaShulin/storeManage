package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"storeManage/models"
	"storeManage/services"
)

type RoleController struct {
	beego.Controller
}

// @Title Role
// @Description 获取菜单
// @Success 200 {object} models.MenuResult "返回结果"
// @Failure 400 {object} models.MenuResult "返回结果"
// @router / [GET]
func (m *RoleController) Role() {
	var reply []models.Role

	reply = services.Role()
	fmt.Println("reply:",reply)
	m.Ctx.Output.Status = 200
	m.Data["json"] = reply
	m.ServeJSON()
}

// @Title Role
// @Description 获取职位
// @Param body body Page true "分页"
// @Success 200 {object} models.GoodsResult "返回结果"
// @Failure 400 {object} models.GoodsResult "返回结果"
// @router /queryAll [POST]
func (r *RoleController) GetRole() {
	//var reply []models.Goods
	var reply models.RoleResult

	in := models.PageInfo{}
	r.ParseForm(&in)

	fmt.Println("PAGE:", in.Page)
	fmt.Println("SIZE:", in.Rows)

	reply = services.GetRoles(in)
	r.Ctx.Output.Status = 200
	r.Data["json"] = reply
	r.ServeJSON()
}

// @Title DeleteRole
// @Description 删除角色
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /remove [POST]
func (r *RoleController) DeleteUser() {
	var reply models.Result

	in := models.IdParam{}
	r.ParseForm(&in)

	reply = services.DeleteRole(in.Id)

	r.Ctx.Output.Status = 200
	reply.Success = true
	reply.Message = "删除用户成功:ID=" + in.Id
	r.Data["json"] = reply
	r.ServeJSON()
}

// @Title SaveRole
// @Description 添加角色
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /save [POST]
func (r *RoleController) SaveRole() {
	var reply models.Result

	in := models.RoleParam{}
	r.ParseForm(&in)

	reply = services.SaveRole(in)

	r.Ctx.Output.Status = 200
	r.Data["json"] = reply
	r.ServeJSON()
}

// @Title SaveEditId
// @Description 获取修改角色ID
// @Param body body models.SaveEditId true "修改商品ID"
// @Success 200 {object} []models.Result "返回结果"
// @Failure 400 {object} []models.Result "返回结果"
// @router /saveEditId [POST]
func (g *RoleController) SaveEditRoleId() {
	var reply models.Result

	in := models.IdParam{}
	g.ParseForm(&in)
	fmt.Println("SaveEditRoleId:",in.Id)
	g.SetSession("roleEditId", in.Id)

	g.Ctx.Output.Status = 200
	reply.Success = true
	reply.Message = "目标角色存储成功:ID=" + in.Id
	g.Data["json"] = reply
	g.ServeJSON()
}

// @Title GetEditRole
// @Description 获取修改商品
// @Success 200 {object} []models.Result "返回结果"
// @Failure 400 {object} []models.Result "返回结果"
// @router /getEditRole [GET]
func (g *RoleController) GetEditRole() {
	var reply models.Role

	roleId := g.GetSession("roleEditId")
	fmt.Print("roleId:", roleId)

	reply = services.GetRole(roleId.(string))

	g.Ctx.Output.Status = 200
	g.Data["json"] = reply
	g.ServeJSON()
}

// @Title EditRole
// @Description 修改角色信息
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /edit [POST]
func (g *RoleController) EditRole() {
	var reply models.Result

	roleId := g.GetSession("roleEditId")
	fmt.Print("roleId:", roleId)

	in := models.RoleParam{}
	g.ParseForm(&in)

	roleInfo :=  make(map[string]interface{})
	roleInfo["NAME"] = in.Name

	reply = services.EditRole(roleId.(string), roleInfo, in.MenuIds)

	g.Ctx.Output.Status = 200
	g.Data["json"] = reply
	g.ServeJSON()
}