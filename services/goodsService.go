package services

import (
	"storeManage/models"
	"storeManage/modelDB"
	"storeManage/common/db"
	"github.com/jinzhu/gorm"
)

func GetGoodss() ([]models.Goods) {

	var result []models.Goods

	db, _ := mysql.GetConn()

	goodss := make([]modelDB.Goods, 0)
	err := db.Raw(modelDB.GET_GOODS_LIST).Scan(&goodss).Error
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

	for _, goods := range goodss {
		var temp models.Goods
		temp.Id = goods.Id
		temp.Name = goods.Name
		temp.Price = goods.Price
		temp.Src = goods.Src
		temp.Remark = goods.Remark
		temp.Number = goods.Number
		temp.Type = goods.Type
		result = append(result, temp)
	}
	return result
}

func GetGoods(id string) (models.Goods) {

	var result models.Goods

	db, _ := mysql.GetConn()

	var goods modelDB.Goods

	sql := modelDB.GET_GOODS_LIST + " WHERE G.ID = ? "
	err := db.Raw(sql,id).Scan(&goods).Error
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

	result.Id = goods.Id
	result.Name = goods.Name
	result.Price = goods.Price
	result.Src = goods.Src
	result.Remark = goods.Remark
	result.Number = goods.Number
	result.Type = goods.Type
	return result
}
