package services

import (
	"storeManage/models"
	"storeManage/modelDB"
	"storeManage/common/db"
	"github.com/jinzhu/gorm"
)

func MenuByUser(userId string) (models.MenuResult) {

	var result models.MenuResult

	db, _ := mysql.GetConn()

	var ids []modelDB.ResultIds
	err := db.Table("TB_USER_ROLE").Where(" USER_ID = ? ", userId).Scan(&ids).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		result.Success = false
		result.Message = "数据库异常"
		return result
	}
	roleIds := make([]string, 0)
	for _, id := range ids {
		roleIds = append(roleIds, id.RoleId)
	}

	err = db.Table("TB_ROLE_MENU").Where(" ROLE_ID IN (?) ", roleIds).Scan(&ids).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		result.Success = false
		result.Message = "数据库异常"
		return result
	}
	menuIds := make([]string, 0)
	for _, id := range ids {
		menuIds = append(menuIds, id.MenuId)
	}


	menus := make([]modelDB.Menu, 0)
	err = db.Table("TB_MENU").Where(" PARENT_ID = ? OR ID IN (?) ", "",menuIds).Scan(&menus).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		result.Success = false
		result.Message = "数据库异常"
		return result
	}
	if err == gorm.ErrRecordNotFound {
		result.Success = false
		result.Message = "当前用户无权限"
		return result
	}

	for _, info := range menus {
		if info.ParentId == "" {
			var menu models.Menu
			menu.Id = info.Id
			menu.Name = info.Name
			menu.IconCls = info.IconCls
			result.Menus = append(result.Menus, menu)
		}
	}
	for _, cInfo := range menus {
		if cInfo.ParentId != "" {
			var cMenu models.Menu
			cMenu.Id = cInfo.Id
			cMenu.Name = cInfo.Name
			cMenu.IconCls = cInfo.IconCls
			cMenu.Href = cInfo.Href
			for i, info := range result.Menus {
				if cInfo.ParentId == info.Id {
					info.Children = append(info.Children, cMenu)
					result.Menus[i] = info
				}
			}
		}
	}

	result.Success = true
	return result
}

func MenuChild() ([]models.Menu) {

	var result []models.Menu

	db, _ := mysql.GetConn()

	var menus []modelDB.Menu
	err := db.Table("TB_MENU").Where(" PARENT_ID != ? ", "").Scan(&menus).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}

	for _, menu := range menus {
		var m models.Menu
		m.Id = menu.Id
		m.Name = menu.Name
		result = append(result, m)
	}

	return result
}
