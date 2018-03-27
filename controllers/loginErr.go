package controllers

import (
	"github.com/astaxie/beego"
)

type LoginErrController struct {
	beego.Controller
}

func (this *LoginErrController) Get() {
	this.TplName = "loginerr.html"
}
func (this *LoginErrController) Post() {
	uname := this.GetString("uname")
	pwd := this.GetString("pwd")

	confUser := beego.AppConfig.String("user")
	confPwd := beego.AppConfig.String("pwd")
	if uname == confUser && pwd == confPwd {
		this.Redirect("/", 302)

	} else {
		this.Redirect("/loginErr", 302)
	}

}
