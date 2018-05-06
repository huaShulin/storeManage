package services

import (
	"storeManage/models"
	"storeManage/modelDB"
	"storeManage/common/db"
	"github.com/jinzhu/gorm"
)

func Menu(userId string) (models.MenuResult) {

	var result models.MenuResult

	db, _ := mysql.GetConn()

	menus := make([]modelDB.Menu, 0)
	err := db.Table("TB_MENU").Scan(&menus).Error
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
