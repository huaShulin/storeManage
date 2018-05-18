package services

import (
	"storeManage/models"
	"storeManage/modelDB"
	"storeManage/common/db"
	"github.com/jinzhu/gorm"
)

func GetRole(info models.PageInfo) (models.RoleResult) {

	var result models.RoleResult

	db, _ := mysql.GetConn()

	roles := make([]modelDB.RoleList, 0)

	sql := modelDB.GET_ROLE_LIST + PageSql(info)
	err := db.Raw(sql).Scan(&roles).Error
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

	for _, role := range roles {
		var temp models.Role
		temp.Id = role.Id
		temp.Name = role.Name
		var number modelDB.ResultNumber
		err = db.Raw("SELECT COUNT(*) AS NUMBER FROM TB_USER_ROLE WHERE ROLE_ID = ? ", role.Id).Scan(&number).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			//result.Success = false
			//result.Message = "数据库异常"
			return result
		}
		temp.Number = number.Number

		var names []modelDB.ResultName
		err = db.Raw("SELECT M.NAME FROM TB_ROLE_MENU RM LEFT JOIN TB_MENU M ON RM.MENU_ID = M.ID WHERE RM.ROLE_ID = ? ", role.Id).Scan(&names).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			//result.Success = false
			//result.Message = "数据库异常"
			return result
		}
		for i, name := range names {
			if i != 0 {
				temp.Menu += "、"
			}
			temp.Menu += name.Name
		}
		result.Roles = append(result.Roles, temp)
	}

	db.Table("TB_USER").Count(&result.Total)

	return result
}

func DeleteRole(id string) models.Result{
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