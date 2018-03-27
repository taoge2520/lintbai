package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsLogin"] = Login
	c.TplName = "index.html"
}

//相关网站
/*
https://wangyasai.github.io/Stars-Emmision/

*/
