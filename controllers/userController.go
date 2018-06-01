package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"storeManage/models"
	"storeManage/services"
	"storeManage/modelDB"
	"math/rand"
	"strconv"
	"storeManage/common/db"
)

type UserController struct {
	beego.Controller
}

// @router /GetHelloWorld
func (u *UserController) GetHelloWorld() {
	fmt.Println("Hello World")
	u.Data["Website"] = "Hello"
	u.Data["Email"] = "World"
	//h.TplName = "index.tpl"
	u.Data["json"] = "Hello World"
	u.ServeJSON()
}

// @Title Login
// @Description 登录
// @Param body body models.Login true "登录结构体"
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /login [POST]
func (u *UserController) Login() {
	var reply models.Result

	fmt.Println("Login")
	login := models.Login{}
	if err := u.ParseForm(&login); err != nil {
		u.Ctx.Output.Status = 200
		reply.Success = false
		reply.Message = "输入参数不合法"
		u.Data["json"] = reply
		u.ServeJSON()
		return
	}

	var user modelDB.User
	reply, user = services.Login(login)

	if reply.Success {
		if user.Status != 1 {
			u.Ctx.Output.Status = 200
			reply.Success = false
			reply.Message = "该用户已离职"
			u.Data["json"] = reply
			u.ServeJSON()
			return
		}
		u.SetSession("user", user)
	}

	u.Ctx.Output.Status = 200
	u.Data["json"] = reply
	u.ServeJSON()
}

// @Title MessageLogin
// @Description 验证码登录
// @Param body body models.Login true "登录结构体"
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /messageLogin [POST]
func (u *UserController) MessageLogin() {
	var reply models.Result

	fmt.Println("Login")
	login := models.Login{}
	if err := u.ParseForm(&login); err != nil {
		u.Ctx.Output.Status = 200
		reply.Success = false
		reply.Message = "输入参数不合法"
		u.Data["json"] = reply
		u.ServeJSON()
		return
	}

	validate := u.GetSession(login.Phone)
	if login.Password != validate.(string) {
		u.Ctx.Output.Status = 200
		reply.Success = false
		reply.Message = "验证码错误"
		u.Data["json"] = reply
		u.ServeJSON()
		return
	}

	var user modelDB.User
	reply, user = services.MessageLogin(login.Phone)

	if reply.Success {
		if user.Status != 1 {
			u.Ctx.Output.Status = 200
			reply.Success = false
			reply.Message = "该用户已离职"
			u.Data["json"] = reply
			u.ServeJSON()
			return
		}
		u.SetSession("user", user)
	}

	u.Ctx.Output.Status = 200
	u.Data["json"] = reply
	u.ServeJSON()
}

// @Title CheckoutLogin
// @Description 退出登录
// @Param body body models.Login true "登录结构体"
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /checkout [GET]
func (u *UserController) CheckoutLogin() {
	//var reply models.Result

	u.DestroySession()


	u.Redirect("/static/back/user/userLogin.html",302)

	/*u.Ctx.Output.Status = 200
	reply.Success = true
	reply.Message = "退出成功"
	u.Data["json"] = reply
	u.ServeJSON()*/
}

// @router /getMsg [POST]
func (u *UserController) GetMsg() {
	var reply models.Result

	send := models.SendMessage{}
	if err := u.ParseForm(&send); err != nil {
		u.Ctx.Output.Status = 200
		reply.Success = false
		reply.Message = "输入参数不合法"
		u.Data["json"] = reply
		u.ServeJSON()
		return
	}

	v1 := rand.Intn(8)
	v2 := rand.Intn(9)
	v3 := rand.Intn(9)
	v4 := rand.Intn(9)
	validate := strconv.Itoa(v1 + 1) + strconv.Itoa(v2) + strconv.Itoa(v3) + strconv.Itoa(v4)
	u.SetSession(send.Phone, validate)
	mysql.Send(send.Phone, validate)
	reply.Success = true
	reply.Message = "验证码已发送"

	u.Ctx.Output.Status = 200
	u.Data["json"] = reply
	u.ServeJSON()
}

// @Title GetLogin
// @Description 获取登录信息
// @Param body body models.Login true "登录结构体"
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /getLogin [POST]
func (u *UserController) GetLogin() {
	var reply models.GetLgin

	fmt.Println("GetLgin")

	user := u.GetSession("user")

	if user == nil {
		u.Ctx.Output.Status = 200
		u.Data["json"] = reply
		u.ServeJSON()
		return
	}

	use := user.(modelDB.User)

	reply.Name = use.Name

	roles := services.GetRolesByUserId(use.Id)
	for i, role := range roles {
		if i != 0 {
			reply.Role += ","
		}
		reply.Role += role.Name
	}

	u.Ctx.Output.Status = 200
	u.Data["json"] = reply
	u.ServeJSON()
}

// @Title User
// @Description 获取用户
// @Param body body Page true "分页"
// @Success 200 {object} models.GoodsResult "返回结果"
// @Failure 400 {object} models.GoodsResult "返回结果"
// @router /queryAll [POST]
func (u *UserController) GetUser() {
	//var reply []models.Goods
	var reply models.UserResult

	in := models.PageInfo{}
	u.ParseForm(&in)

	fmt.Println("PAGE:", in.Page)
	fmt.Println("SIZE:", in.Rows)

	reply = services.GetUsers(in)
	u.Ctx.Output.Status = 200
	u.Data["json"] = reply
	u.ServeJSON()
}

// @Title StatusUser
// @Description 修改用户状态
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /status [POST]
func (u *UserController) StatusUser() {
	var reply models.Result

	in := models.IdParam{}
	u.ParseForm(&in)

	reply = services.StatusUser(in.Id)

	u.Ctx.Output.Status = 200
	reply.Success = true
	reply.Message = "修改用户状态成功:ID=" + in.Id
	u.Data["json"] = reply
	u.ServeJSON()
}

// @Title DeleteUser
// @Description 删除用户
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /remove [POST]
func (u *UserController) DeleteUser() {
	var reply models.Result

	in := models.IdParam{}
	u.ParseForm(&in)

	reply = services.DeleteUser(in.Id)

	u.Ctx.Output.Status = 200
	reply.Success = true
	reply.Message = "删除用户成功:ID=" + in.Id
	u.Data["json"] = reply
	u.ServeJSON()
}

// @Title SaveEditId
// @Description 获取修改角色ID
// @Param body body models.SaveEditId true "修改商品ID"
// @Success 200 {object} []models.Result "返回结果"
// @Failure 400 {object} []models.Result "返回结果"
// @router /saveEditId [POST]
func (g *UserController) SaveEditUserId() {
	var reply models.Result

	in := models.IdParam{}
	g.ParseForm(&in)
	fmt.Println("SaveEditUserId:",in.Id)
	g.SetSession("userEditId", in.Id)

	g.Ctx.Output.Status = 200
	reply.Success = true
	reply.Message = "目标用户存储成功:ID=" + in.Id
	g.Data["json"] = reply
	g.ServeJSON()
}

// @Title GetEditRole
// @Description 获取修改商品
// @Success 200 {object} []models.Result "返回结果"
// @Failure 400 {object} []models.Result "返回结果"
// @router /getEditUser [GET]
func (g *UserController) GetEditUser() {
	var reply models.User

	userId := g.GetSession("userEditId")
	fmt.Print("userId:", userId)

	reply = services.GetUser(userId.(string))

	g.Ctx.Output.Status = 200
	g.Data["json"] = reply
	g.ServeJSON()
}

// @Title EditRole
// @Description 修改角色信息
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /edit [POST]
func (g *UserController) EditUser() {
	var reply models.Result

	userId := g.GetSession("userEditId")
	fmt.Print("userId:", userId)

	in := models.UserParam{}
	g.ParseForm(&in)

	roleInfo :=  make(map[string]interface{})
	roleInfo["NAME"] = in.Name
	roleInfo["PHONE"] = in.Phone

	reply = services.EditUser(userId.(string), roleInfo, in.RoleIds)

	g.Ctx.Output.Status = 200
	g.Data["json"] = reply
	g.ServeJSON()
}

// @Title SaveUser
// @Description 添加用户
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /save [POST]
func (r *UserController) SaveUser() {
	var reply models.Result

	in := models.UserParam{}
	r.ParseForm(&in)

	reply = services.SaveUser(in)

	r.Ctx.Output.Status = 200
	r.Data["json"] = reply
	r.ServeJSON()
}

// @Title Me
// @Description 获取用户
// @Param body body Page true "分页"
// @Success 200 {object} models.GoodsResult "返回结果"
// @Failure 400 {object} models.GoodsResult "返回结果"
// @router /queryMe [POST]
func (u *UserController) Me() {
	//var reply []models.Goods
	var reply []models.User

	user := u.GetSession("user")
	use := user.(modelDB.User)
	fmt.Print("获取个人信息",use)

	temp := services.GetMe(use.Id)
	reply = append(reply, temp)
	u.Ctx.Output.Status = 200
	u.Data["json"] = reply
	u.ServeJSON()
}

// @router /getEditMe [GET]
func (u *UserController) GetMe() {
	var reply models.User

	user := u.GetSession("user")
	use := user.(modelDB.User)
	fmt.Print("获取菜单",use)

	reply = services.GetMe(use.Id)

	v1 := rand.Intn(8)
	v2 := rand.Intn(9)
	v3 := rand.Intn(9)
	v4 := rand.Intn(9)
	validate := strconv.Itoa(v1 + 1) + strconv.Itoa(v2) + strconv.Itoa(v3) + strconv.Itoa(v4)
	u.SetSession("validate", validate)
	mysql.Send(reply.Phone, validate)

	u.Ctx.Output.Status = 200
	u.Data["json"] = reply
	u.ServeJSON()
}

// @Title EditRole
// @Description 修改角色信息
// @Success 200 {object} models.Result "返回结果"
// @Failure 400 {object} models.Result "返回结果"
// @router /editMe [POST]
func (g *UserController) EditMe() {
	var reply models.Result

	in := models.MeParam{}
	g.ParseForm(&in)

	validate := g.GetSession("validate")
	if in.Validate != validate.(string) {
		g.Ctx.Output.Status = 200
		reply.Success = false
		reply.Message = "验证码错误"
		g.Data["json"] = reply
		g.ServeJSON()
		return
	}

	roleInfo :=  make(map[string]interface{})
	roleInfo["PHONE"] = in.Phone
	if in.Password != "" {
		roleInfo["PASSWORD"] = mysql.GetMd5String(in.Password)
	}

	user := g.GetSession("user")
	use := user.(modelDB.User)

	reply = services.EditMe(use.Id, roleInfo)

	g.Ctx.Output.Status = 200
	g.Data["json"] = reply
	g.ServeJSON()
}