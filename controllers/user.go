package controllers

import (
	"MyBlog/models/class"
	"github.com/russross/blackfriday"
)

type UserController struct {
	BaseController
}

func (c *UserController) Profile() {

	id := c.Ctx.Input.Param(":id")
	u := &class.User{Id: id}
	u.ReadDB()
	c.Data["u"] = u
	as := class.Article{Author: u}.Gets()
	for id,a:=range as{
		unsafe := blackfriday.MarkdownBasic([] byte(a.Content))
		//html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
		a.Content=string(unsafe);
		as[id]=a;
	}
	replys := class.Reply{Author: u}.Gets()
	c.Data["articles"] = as
	c.Data["replys"] = replys
	c.TplName = "user/profile.html"
}
func (c *UserController) LoginJoin() {
	c.TplName = "user/login.html"
}
func  (c *UserController)RegisterPage()  {
	c.TplName="user/register.html"
}
func (c *UserController) PageSetting() {
	c.CheckLogin()
	c.TplName = "user/setting.html"
}
func (c *UserController) Markdown()  {
	c.TplName="markdown.html"
}