package controllers

import (
	"github.com/astaxie/beego"
	"github.com/zhengjianwen/hr_permission/models"
)

type IndexController struct {
	beego.Controller
}

func (index *IndexController)Get()  {
	user := &models.User{Name:"张三",Position:"管理员"}
	index.Data["user"] = user

	index.TplName = "index.html"
}

func (index *IndexController)Ifrom()  {
	index.TplName = "index4.html"
}