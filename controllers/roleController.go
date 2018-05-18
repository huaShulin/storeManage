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
// @Description 获取职位
// @Param body body Page true "分页"
// @Success 200 {object} models.GoodsResult "返回结果"
// @Failure 400 {object} models.GoodsResult "返回结果"
// @router /queryAll [POST]
func (r *RoleController) GetUser() {
	//var reply []models.Goods
	var reply models.RoleResult

	in := models.PageInfo{}
	r.ParseForm(&in)

	fmt.Println("PAGE:", in.Page)
	fmt.Println("SIZE:", in.Rows)

	reply = services.GetRole(in)
	r.Ctx.Output.Status = 200
	r.Data["json"] = reply
	r.ServeJSON()
}

// @Title StatusUser
// @Description 修改用户状态
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /status [POST]
func (r *RoleController) StatusUser() {
	var reply models.Result

	in := models.IdParam{}
	r.ParseForm(&in)

	reply = services.StatusUser(in.Id)

	r.Ctx.Output.Status = 200
	reply.Success = true
	reply.Message = "修改用户状态成功:ID=" + in.Id
	r.Data["json"] = reply
	r.ServeJSON()
}

// @Title DeleteUser
// @Description 删除用户
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /remove [POST]
func (r *RoleController) DeleteUser() {
	var reply models.Result

	in := models.IdParam{}
	r.ParseForm(&in)

	reply = services.DeleteUser(in.Id)

	r.Ctx.Output.Status = 200
	reply.Success = true
	reply.Message = "删除用户成功:ID=" + in.Id
	r.Data["json"] = reply
	r.ServeJSON()
}