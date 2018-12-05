package routers

import (
	"github.com/zhengjianwen/hr_permission/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.LoginController{})
    beego.Router("/index.html", &controllers.IndexController{})
    beego.Router("/index4.html", &controllers.IndexController{},"*:Ifrom")
}
