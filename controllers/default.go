package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	clientIP := c.Ctx.Input.IP()
	fmt.Println(clientIP)
	c.Data["ip"] = clientIP
	c.TplName = "index.tpl"
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
