package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {

	this.Data["IsLogin"] = Login

	this.TplName = "home.html"
}
