package router

import (
	"github.com/astaxie/beego"
	"webadmin/controller/AppController"
)

func init() {
	beego.Router("/index", &AppController.AppController{}, "*:AppList")

}
