package routers

import (
	"MyBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.UserController{}, `get:LoginJoin;post:Login`)
	beego.Router("/register", &controllers.UserController{}, `get:RegisterPage;post:Register`)
	beego.Router("/logout", &controllers.UserController{}, `get:Logout`)
	beego.Router("/setting", &controllers.UserController{}, `get:PageSetting;post:Setting`)
	beego.Router("/user/:id", &controllers.UserController{}, `get:Profile`)
	beego.Router("/article/new", &controllers.ArticleController{}, `get:PageNew;post:New`)
	beego.Router("/article/:id([0-9]+)", &controllers.ArticleController{}, `get:Get`)
	beego.Router("/article/delete/:id([0-9]+)", &controllers.ArticleController{}, `get:Del`)
	beego.Router("/article/edit/:id([0-9]+)", &controllers.ArticleController{}, `get:PageEdit;post:Edit`)
	beego.Router("/archive", &controllers.ArticleController{}, `get:Archive`)
	beego.Router("/reply/new", &controllers.ReplyController{}, `post:New`)
	beego.Router("/user/profile", &controllers.UserController{}, `get:Profile`)
	beego.Router("/markdown", &controllers.UserController{}, `get:Markdown`)
}
