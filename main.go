package main

import (
	_ "storeManage/routers"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		beego.BConfig.WebConfig.Session.SessionOn = true
	}
	beego.Run()
}

