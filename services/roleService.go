package services

import (
	"storeManage/models"
	"storeManage/modelDB"
	"storeManage/common/db"
	"github.com/jinzhu/gorm"
)

func GetRoles(info models.PageInfo) (models.RoleResult) {

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

	db.Table("TB_ROLE").Count(&result.Total)

	return result
}

func DeleteRole(id string) models.Result{
	var result models.Result

	db, _ := mysql.GetConn()
	db.Begin()

	err := db.Table("TB_ROLE").Where(" ID = ? ", id).Delete(&result).Error
	if err != nil {
		db.Rollback()
		result.Success = false
		result.Message = "删除失败"
		return result
	}

	err = db.Table("TB_ROLE_MENU").Where(" ROLE_ID = ? ", id).Delete(&result).Error
	if err != nil {
		db.Rollback()
		result.Success = false
		result.Message = "删除失败"
		return result
	}

	db.Commit()
	result.Success = true
	result.Message = "删除成功"
	return result
}

func SaveRole(info models.RoleParam) models.Result{
	var result models.Result

	var role modelDB.SaveRole
	role.Id = mysql.GetId()
	role.Name = info.Name

	db, _ := mysql.GetConn()
	err := db.Table("TB_ROLE").Save(&role).Error
	if err != nil {
		result.Success = false
		result.Message = "保存失败"
		return result
	}

	for _, menuId := range info.MenuIds {
		var roleMenu modelDB.SaveRoleMenu
		roleMenu.Id = mysql.GetId()
		roleMenu.RoleId = role.Id
		roleMenu.MenuId = menuId
		err = db.Table("TB_ROLE_MENU").Save(&roleMenu).Error
		if err != nil {
			result.Success = false
			result.Message = "保存失败"
			return result
		}
	}

	result.Success = true
	result.Message = "保存成功"
	return result
}

func GetRole(id string) (models.Role) {

	var result models.Role

	db, _ := mysql.GetConn()

	var role modelDB.RoleList

	//sql := modelDB.GET_ROLE_LIST + " WHERE ID = ? "
	//err := db.Raw(sql,id).Scan(&goods).Error
	err := db.Table("TB_ROLE").Where(" ID = ? ",id).Scan(&role).Error
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

	result.Id = role.Id
	result.Name = role.Name
	return result
}

func EditRole(id string, roleInfo map[string]interface{}, menuIds []string) models.Result{
	var result models.Result

	db, _ := mysql.GetConn()
	db.Begin()

	err := db.Table("TB_ROLE").Where(" ID = ? ", id).Update(roleInfo).Error
	if err != nil {
		db.Rollback()
		result.Success = false
		result.Message = "修改失败"
		return result
	}

	err = db.Table("TB_ROLE_MENU").Where(" ROLE_ID = ? ", id).Delete(&result).Error
	if err != nil {
		db.Rollback()
		result.Success = false
		result.Message = "修改失败"
		return result
	}

	for _, menuId := range menuIds {
		var roleMenu modelDB.SaveRoleMenu
		roleMenu.Id = mysql.GetId()
		roleMenu.RoleId = id
		roleMenu.MenuId = menuId
		err = db.Table("TB_ROLE_MENU").Save(&roleMenu).Error
		if err != nil {
			db.Rollback()
			result.Success = false
			result.Message = "保存失败"
			return result
		}
	}


	result.Success = true
	result.Message = "修改成功"
	db.Commit()
	return result
}

func Role() ([]models.Role) {

	var result []models.Role

	db, _ := mysql.GetConn()

	var roles []modelDB.RoleList
	err := db.Table("TB_ROLE").Scan(&roles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}

	for _, role := range roles {
		var r models.Role
		r.Id = role.Id
		r.Name = role.Name
		result = append(result, r)
	}

	return result
}

func GetRolesByUserId(id string) ([]modelDB.RoleList) {

	var result []modelDB.RoleList

	db, _ := mysql.GetConn()

	var userRoles []modelDB.ResultIds
	err := db.Table("TB_USER_ROLE").Where(" USER_ID = ? ", id).Scan(&userRoles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return result
	}
	var ids []string
	for _, userRole := range userRoles {
		ids = append(ids, userRole.RoleId)
	}

	err = db.Table("TB_ROLE").Where(" ID IN (?) ", ids).Scan(&result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return result
	}

	return result
}