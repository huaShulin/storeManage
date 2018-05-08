package services

import (
	"storeManage/models"
	"storeManage/modelDB"
	"storeManage/common/db"
	"github.com/jinzhu/gorm"
)

func GetGoodss(info models.PageInfo) ([]models.Goods) {

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
	result.Purchase = goods.Purchase
	result.Type = goods.Type
	return result
}

func EditGoods(id string, goodsInfo map[string]interface{}) models.Result{
	var result models.Result

	db, _ := mysql.GetConn()
	err := db.Table("TB_GOODS").Where(" ID = ? ", id).Update(goodsInfo).Error
	if err != nil {
		result.Success = false
		result.Message = "修改失败"
	}

	result.Success = true
	result.Message = "修改成功"
	return result
}

func SaveGoods(info models.GoodsParam) models.Result{
	var result models.Result

	var goods modelDB.SaveGoods
	goods.Id = mysql.GetId()
	goods.Name = info.Name
	goods.Price = info.Price
	goods.Number = info.Number
	goods.Purchase = info.Purchase
	goods.Remark = info.Remark
	goods.TypeId = info.TypeId

	db, _ := mysql.GetConn()
	err := db.Table("TB_GOODS").Save(&goods).Error
	if err != nil {
		result.Success = false
		result.Message = "保存失败"
	}

	result.Success = true
	result.Message = "保存成功"
	return result
}

func DeleteGoods(id string) models.Result{
	var result models.Result

	db, _ := mysql.GetConn()
	err := db.Table("TB_GOODS").Where(" ID = ? ", id).Delete(&result).Error
	if err != nil {
		result.Success = false
		result.Message = "删除失败"
	}

	result.Success = true
	result.Message = "删除成功"
	return result
}