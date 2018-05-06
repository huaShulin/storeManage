package services

import (
	"storeManage/models"
	"storeManage/modelDB"
	"storeManage/common/db"
	"github.com/jinzhu/gorm"
)

func Login(info models.Login) (models.Result, string) {

	var result models.Result

	db, _ := mysql.GetConn()

	var user modelDB.User
	err := db.Table("TB_USER").Where(" USERNAME = ? ", info.Username).Scan(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		result.Success = false
		result.Message = "数据库异常"
		return result, ""
	}
	if err == gorm.ErrRecordNotFound {
		result.Success = false
		result.Message = "用户不存在"
		return result, ""
	}
	if mysql.GetMd5String(info.Password) != user.Password {
		result.Success = false
		result.Message = "密码错误"
		return result, ""
	}

	//存入session
	//l.SetSession("user", user.Id)
	//l.Data["Website"] = user.Id

	result.Success = true
	result.Message = "登陆成功"
	return result, user.Id
}
