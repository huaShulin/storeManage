package mysql

import (
	"io"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/base64"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//获取数据库连接
func GetConn() (*gorm.DB, error) {
	username := "root"
	password := "root"
	dburl := "@(39.106.96.217:3306)/STORE_MANAGE?charset=utf8&parseTime=True&loc=Local"
	str := username + ":" + password + dburl
	db, err := gorm.Open("mysql", str)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
	return db, nil
}

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func GetId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}