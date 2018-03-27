package controllers

import (
	"github.com/astaxie/beego"
)

type ToolsController struct {
	beego.Controller
}

func (this *ToolsController) Get() {

	this.TplName = "tools.html"

}
