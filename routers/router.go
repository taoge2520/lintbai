package routers

import (
	"lintbai/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/home", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/loginErr", &controllers.LoginErrController{})
	beego.Router("/tools", &controllers.ToolsController{})
	beego.Router("/photo", &controllers.PhotoController{})
	beego.Router("/dh", &controllers.DhController{})
}
