package controllers

import (
	"github.com/astaxie/beego"
	_"github.com/astaxie/beego/orm"
	_"server/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "Index.html"
}