package services

import (
	"storeManage/models"
	"storeManage/modelDB"
	"storeManage/common/db"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

func GetOrder(info models.PageInfo) models.OrderResult {

	var result models.OrderResult

	db, _ := mysql.GetConn()

	sql := modelDB.GET_ORDER_LIST + PageSql(info)
	orders := make([]modelDB.Order, 0)
	err := db.Raw(sql).Scan(&orders).Error
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

	for _, order := range orders {
		var temp models.Order
		temp.Id = order.Id
		temp.UserId = order.UserId
		temp.UserName = order.UserName
		temp.Sum = order.Sum
		//temp.Time = order.Time
		t, _ := time.ParseInLocation("20060102150405", order.Time, time.Local)
		temp.Time = t.Format("2006-01-02 15:04:05")
		temp.Remark = order.Remark

		var orderGoodss []modelDB.OrderGoods
		db.Table("TB_ORDER_GOODS").Where(" ORDER_ID = ? ",order.Id).Scan(&orderGoodss)
		for i, orderGoods := range orderGoodss {
			if i != 0 {
				temp.Details += "；"
			}
			temp.Details += "商品名：" + orderGoods.Name + "，价格：" + strconv.FormatFloat(orderGoods.Price, 'f', -1, 64) +
				"元，数量：" + strconv.Itoa(orderGoods.Number) + "，商品总价：" + strconv.FormatFloat(orderGoods.Sum, 'f', -1, 64) + "元"
		}

		result.Order = append(result.Order, temp)

		/*
		goodsInfo :=  make(map[string]interface{})
		goodsInfo["NUMBER"] = goods.Number - temp.Number
		*/
	}

	db.Table("TB_ORDER").Count(&result.Total)

	return result
}

func GetMeOrder(info models.PageInfo, userId string) models.OrderResult {

	var result models.OrderResult

	db, _ := mysql.GetConn()

	sql := modelDB.GET_ORDER_LIST + " WHERE USER_ID = ? " + PageSql(info)
	orders := make([]modelDB.Order, 0)
	err := db.Raw(sql, userId).Scan(&orders).Error
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

	for _, order := range orders {
		var temp models.Order
		temp.Id = order.Id
		temp.UserId = order.UserId
		temp.UserName = order.UserName
		temp.Sum = order.Sum
		//temp.Time = order.Time
		t, _ := time.ParseInLocation("20060102150405", order.Time, time.Local)
		temp.Time = t.Format("2006-01-02 15:04:05")
		temp.Remark = order.Remark

		var orderGoodss []modelDB.OrderGoods
		db.Table("TB_ORDER_GOODS").Where(" ORDER_ID = ? ",order.Id).Scan(&orderGoodss)
		for i, orderGoods := range orderGoodss {
			if i != 0 {
				temp.Details += "；"
			}
			temp.Details += "商品名：" + orderGoods.Name + "，价格：" + strconv.FormatFloat(orderGoods.Price, 'f', -1, 64) +
				"元，数量：" + strconv.Itoa(orderGoods.Number) + "，商品总价：" + strconv.FormatFloat(orderGoods.Sum, 'f', -1, 64) + "元"
		}

		result.Order = append(result.Order, temp)

		/*
		goodsInfo :=  make(map[string]interface{})
		goodsInfo["NUMBER"] = goods.Number - temp.Number
		*/
	}

	db.Table("TB_ORDER").Where(" USER_ID = ? ", userId).Count(&result.Total)

	return result
}

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

func CreateOrder(userId string, orderGoodss map[string]int, remark string) (models.Result) {
	var result models.Result

	db, _ := mysql.GetConn()
	db.Begin()

	orderId := mysql.GetId()

	ids := make([]string, 0)
	for id, _ := range orderGoodss {
		ids = append(ids, id)
	}

	goodss := make([]modelDB.Goods, 0)
	sql := modelDB.GET_GOODS_LIST + " WHERE G.ID IN (?) "
	db.Raw(sql, ids).Scan(&goodss)

	var allSum float64
	for _, goods := range goodss {
		//减去商品库存
		goodsMap := make(map[string]interface{})
		number := goods.Number - orderGoodss[goods.Id]
		if number < 0 {
			db.Rollback()
			result.Success = false
			result.Message = "商品数量不足"
			return result
		}
		goodsMap["NUMBER"] = number

		db.Table("TB_GOODS").Where(" ID = ? ", goods.Id).Update(goodsMap)

		var tempOrderGoods modelDB.OrderGoods
		tempOrderGoods.Id = mysql.GetId()
		tempOrderGoods.Name = goods.Name
		tempOrderGoods.Price = goods.Price
		tempOrderGoods.Number = orderGoodss[goods.Id]
		tempOrderGoods.Type = goods.Type
		tempOrderGoods.Sum = float64(tempOrderGoods.Number) * tempOrderGoods.Price
		tempOrderGoods.GoodsId = goods.Id
		tempOrderGoods.OrderId = orderId
		db.Table("TB_ORDER_GOODS").Save(&tempOrderGoods)

		allSum += tempOrderGoods.Sum
	}

	var tempOrder modelDB.Order
	tempOrder.Id = orderId
	tempOrder.UserId = userId
	var tempUser modelDB.User
	db.Table("TB_USER").Where(" ID = ? ", userId).Scan(&tempUser)
	tempOrder.UserName = tempUser.Name
	tempOrder.Sum = allSum
	tempOrder.Time = time.Now().Format("20060102150405")
	tempOrder.Remark = remark
	db.Table("TB_ORDER").Save(&tempOrder)

	db.Commit()
	result.Success = true
	result.Message = "成功"
	return result
}

func GetOrderSum(orderGoodss map[string]int) (models.Result) {
	var result models.Result

	db, _ := mysql.GetConn()

	ids := make([]string, 0)
	for id, _ := range orderGoodss {
		ids = append(ids, id)
	}

	goodss := make([]modelDB.Goods, 0)
	sql := modelDB.GET_GOODS_LIST + " WHERE G.ID IN (?) "
	db.Raw(sql, ids).Scan(&goodss)

	var sum float64
	for _, goods := range goodss {
		sum += goods.Price * float64(orderGoodss[goods.Id])
	}

	result.Success = true
	result.Message = strconv.FormatFloat(sum, 'f', -1, 64)
	return result
}