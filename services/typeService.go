package services

import (
	"storeManage/models"
	"storeManage/common/db"
	"github.com/jinzhu/gorm"
	"storeManage/modelDB"
)

func Type() ([]models.Type) {

	var result []models.Type

	db, _ := mysql.GetConn()

	types := make([]modelDB.Type, 0)
	err := db.Table("TB_TYPE").Scan(&types).Error
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

	for _, ty := range types {
		var temp models.Type
		temp.Id = ty.Id
		temp.Name = ty.Name
		temp.ParentId = ty.ParentId
		result = append(result, temp)
	}

	return result
}
