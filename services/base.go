package services

import (
	"storeManage/models"
	"strconv"
)

func PageSql(info models.PageInfo) string {
	//PAGE = " ORDER BY ID LIMIT ?,? "
	sql := " ORDER BY ID LIMIT " + strconv.Itoa((info.Page-1) * info.Rows) + "," + strconv.Itoa((info.Page) * info.Rows) + " "
	return sql
}