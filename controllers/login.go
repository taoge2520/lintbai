package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	Login = false
	if this.Input().Get("exit") == "true" {
		this.Redirect("/", 302)
	}

	this.TplName = "login.html"
}

var Login bool

func (this *LoginController) Post() {
	uname := this.GetString("uname")
	pwd := this.GetString("pwd")

	confUser := beego.AppConfig.String("user")
	confPwd := beego.AppConfig.String("pwd")
	if uname == confUser && pwd == confPwd {
		Login = true
		this.Data["IsLogin"] = Login
		this.Redirect("/", 302)

	} else {
		this.Redirect("/loginErr", 302)
	}

}
