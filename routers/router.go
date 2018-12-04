package routers

import (
	"github.com/zhengjianwen/hr_permission/hr_permission/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
