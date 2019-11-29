package AppController

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type AppController struct {
	beego.Controller
}

func (p *AppController) AppList() {

	logs.Debug("enter index controller")

	p.Layout = "layout/layout.html"
	appList, err := model.GetAllAppInfo()
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("服务器繁忙")
		p.TplName = "app/error.html"

		logs.Warn("get app list failed, err:%v", err)
		return
	}

	logs.Debug("get app list succ, data:%v", appList)
	p.Data["applist"] = appList

	p.TplName = "app/index.html"
}
