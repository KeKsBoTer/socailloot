package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	AuthController
}

func (c *IndexController) Get() {
	choice := Choice(c.Ctx.Input.Param(":choice"))
	if len(choice) < 1 {
		choice = Hot
	}
	if !choice.IsValid() {
		c.Abort("404")
		return
	}
	posts, err := getPostsForTopic(c.User, nil, choice)
	if err != nil {
		beego.Error(err)
		c.Abort("500")
		return
	}
	c.Data["Posts"] = posts
	c.Data["Choice"] = choice

	c.Data["Title"] = "Socailloot: like reddit but different"
	c.TplName = "pages/index.tpl"
}
