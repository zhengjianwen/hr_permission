package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["Name"] = "海瑞权限管理系统"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "login2.html"
}
