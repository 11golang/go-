package main

import (
	"github.com/astaxie/beego"
	_ "webbee/routers"
	"webbee/utils"
)

func main() {
	utils.InitMysql()
	beego.Run()
}
