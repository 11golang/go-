package routers

import (
	"github.com/astaxie/beego"
	"webbee/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/new2.html", &controllers.MainController{}, "GET:Getinfo;POST:Postinfo")
	beego.Router("/GrabdataDisat.html",&controllers.GrabdataDisat{},"GET:Getgradbta")
	beego.Router("/bottle.html",&controllers.LogController{},"GET:Log;POST:Logpost")
}
