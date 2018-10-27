package routers

import (
	"github.com/astaxie/beego"
	"storeserv/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.RegionController{})
	beego.AutoRouter(&controllers.StoresController{})
	beego.AutoRouter(&controllers.LoginController{})
}
