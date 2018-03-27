package controllers

import (
	"github.com/astaxie/beego"
)

type DhController struct {
	beego.Controller
}

func (c *DhController) Get() {

	c.TplName = "navigation.html"
}
