package services

import (
	"storeManage/models"
	"storeManage/modelDB"
	"storeManage/common/db"
	"github.com/jinzhu/gorm"
)

func Login(info models.Login) (models.Result, modelDB.User) {

	var result models.Result

	db, _ := mysql.GetConn()

	var user modelDB.User
	err := db.Table("TB_USER").Where(" PHONE = ? ", info.Phone).Scan(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		result.Success = false
		result.Message = "数据库异常"
		return result, user
	}
	if err == gorm.ErrRecordNotFound {
		result.Success = false
		result.Message = "用户不存在"
		return result, user
	}
	if mysql.GetMd5String(info.Password) != user.Password {
		result.Success = false
		result.Message = "密码错误"
		return result, user
	}

	//存入session
	//l.SetSession("user", user.Id)
	//l.Data["Website"] = user.Id

	result.Success = true
	result.Message = "登陆成功"
	return result, user
}


func GetUsers(info models.PageInfo) (models.UserResult) {

	var result models.UserResult

	db, _ := mysql.GetConn()

	users := make([]modelDB.UserList, 0)

	sql := modelDB.GET_USER_LIST + PageSql(info)
	err := db.Raw(sql).Scan(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		//result.Success = false
		//result.Message = "数据库异常"
		return result
	}
	if err == gorm.ErrRecordNotFound {
		//result.Success = false
		//result.Message = "当前用户无权限"
		return result
	}

	for _, user := range users {
		var temp models.User
		temp.Id = user.Id
		temp.Name = user.Name
		temp.Phone = user.Phone
		if user.Status == 1 {
			temp.Status = "在职"
		} else {
			temp.Status = "离职"
		}

		var userRoles []modelDB.ResultIds
		err := db.Table("TB_USER_ROLE").Where(" USER_ID = ? ", user.Id).Scan(&userRoles).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return result
		}
		var ids []string
		for _, userRole := range userRoles {
			ids = append(ids, userRole.RoleId)
		}
		var roles []modelDB.RoleList
		err = db.Table("TB_ROLE").Where(" ID IN (?) ", ids).Scan(&roles).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return result
		}
		for i, role := range roles {
			if i != 0 {
				temp.Role += ","
			}
			temp.Role += role.Name
		}
		result.Users = append(result.Users, temp)
	}

	db.Table("TB_USER").Count(&result.Total)

	return result
}

func StatusUser(id string) models.Result{
	var result models.Result

	db, _ := mysql.GetConn()
	var user modelDB.User
	err := db.Table("TB_USER").Where(" ID = ? ", id).Scan(&user).Error
	if err != nil {
		result.Success = false
		result.Message = "修改状态读取当前状态失败"
		return result
	}

	userInfo :=  make(map[string]interface{})

	if user.Status == 1 {
		userInfo["STATUS"] = 2
	} else {
		userInfo["STATUS"] = 1
	}
	err = db.Table("TB_USER").Where(" ID = ? ", id).Update(userInfo).Error
	if err != nil {
		result.Success = false
		result.Message = "修改状态失败"
		return result
	}

	result.Success = true
	result.Message = "修改状态成功"
	return result
}

func DeleteUser(id string) models.Result{
	var result models.Result

	db, _ := mysql.GetConn()
	err := db.Table("TB_USER").Where(" ID = ? ", id).Delete(&result).Error
	if err != nil {
		result.Success = false
		result.Message = "删除失败"
		return result
	}

	result.Success = true
	result.Message = "删除成功"
	return result
}

func GetUser(id string) (models.User) {

	var result models.User

	db, _ := mysql.GetConn()

	var user modelDB.User
	err := db.Table("TB_USER").Where(" ID = ? ",id).Scan(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		//result.Success = false
		//result.Message = "数据库异常"
		return result
	}
	if err == gorm.ErrRecordNotFound {
		//result.Success = false
		//result.Message = "当前用户无权限"
		return result
	}

	result.Id = user.Id
	result.Name = user.Name
	result.Phone = user.Phone
	return result
}

func EditUser(id string, roleInfo map[string]interface{}, roleIds []string) models.Result{
	var result models.Result

	db, _ := mysql.GetConn()
	db.Begin()

	err := db.Table("TB_USER").Where(" ID = ? ", id).Update(roleInfo).Error
	if err != nil {
		db.Rollback()
		result.Success = false
		result.Message = "修改失败"
		return result
	}

	err = db.Table("TB_USER_ROLE").Where(" USER_ID = ? ", id).Delete(&result).Error
	if err != nil {
		db.Rollback()
		result.Success = false
		result.Message = "修改失败"
		return result
	}

	for _, roleId := range roleIds {
		var userRole modelDB.SaveUserRole
		userRole.Id = mysql.GetId()
		userRole.UserId = id
		userRole.RoleId = roleId
		err = db.Table("TB_USER_ROLE").Save(&userRole).Error
		if err != nil {
			db.Rollback()
			result.Success = false
			result.Message = "修改失败"
			return result
		}
	}

	result.Success = true
	result.Message = "修改成功"
	db.Commit()
	return result
}