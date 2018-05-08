package services

import (
	"storeManage/models"
	"storeManage/modelDB"
	"storeManage/common/db"
	"github.com/jinzhu/gorm"
)

func GetNowOrder(info models.PageInfo,orderGoodss map[string]int) (models.OrderGoodsResult) {

	var result models.OrderGoodsResult

	db, _ := mysql.GetConn()

	ids := make([]string, 0)
	for id, _ := range orderGoodss {
		ids = append(ids, id)
	}

	sql := modelDB.GET_GOODS_LIST + " WHERE G.ID IN (?) "+ PageSql(info)
	goodss := make([]modelDB.Goods, 0)
	err := db.Raw(sql, ids).Scan(&goodss).Error
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
		var temp models.OrderGoods
		temp.Id = goods.Id
		temp.Name = goods.Name
		temp.Price = goods.Price
		temp.OrderNumber = orderGoodss[goods.Id]
		temp.Number = goods.Number
		temp.Type = goods.Type
		result.OrderGoods = append(result.OrderGoods, temp)

		/*
		goodsInfo :=  make(map[string]interface{})
		goodsInfo["NUMBER"] = goods.Number - temp.Number
		*/
	}

	db.Table("TB_GOODS").Where(" ID IN (?) ", ids).Count(&result.Total)

	return result
}