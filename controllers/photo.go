package controllers

import (
	"github.com/astaxie/beego"
)

type PhotoController struct {
	beego.Controller
}

func (this *PhotoController) Get() {
	this.TplName = "photo.html"
}
