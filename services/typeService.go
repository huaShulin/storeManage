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

func GetTypes(info models.PageInfo) models.TypePageResult {

	var result models.TypePageResult

	db, _ := mysql.GetConn()

	types := make([]modelDB.TypePage, 0)

	sql := modelDB.GET_TYPE_LIST + PageSql(info)
	err := db.Raw(sql).Scan(&types).Error
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

	for _, typ := range types {
		var temp models.TypePage
		temp.Id = typ.Id
		temp.Name = typ.Name
		temp.Parent = typ.Parent
		result.TypePages = append(result.TypePages, temp)
	}

	db.Table("TB_TYPE").Count(&result.Total)

	return result
}

func GetTypeById(id string) (models.Type) {

	var result models.Type

	db, _ := mysql.GetConn()

	var typ modelDB.Type

	err := db.Table("TB_TYPE").Where(" ID = ? ",id).Scan(&typ).Error
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

	result.Id = typ.Id
	result.Name = typ.Name
	return result
}

func EditType(id string, typeInfo map[string]interface{}) models.Result{
	var result models.Result

	db, _ := mysql.GetConn()
	err := db.Table("TB_TYPE").Where(" ID = ? ", id).Update(typeInfo).Error
	if err != nil {
		result.Success = false
		result.Message = "修改失败"
		return result
	}

	result.Success = true
	result.Message = "修改成功"
	return result
}

func SaveType(info models.GoodsParam) models.Result{
	var result models.Result

	var typ modelDB.Type
	typ.Id = mysql.GetId()
	typ.Name = info.Name
	typ.ParentId = info.TypeId

	db, _ := mysql.GetConn()
	err := db.Table("TB_TYPE").Save(&typ).Error
	if err != nil {
		result.Success = false
		result.Message = "保存失败"
		return result
	}

	result.Success = true
	result.Message = "保存成功"
	return result
}

func DeleteType(id string) models.Result{
	var result models.Result

	db, _ := mysql.GetConn()
	err := db.Table("TB_TYPE").Where(" ID = ? ", id).Delete(&result).Error
	if err != nil {
		result.Success = false
		result.Message = "删除失败"
		return result
	}

	result.Success = true
	result.Message = "删除成功"
	return result
}
