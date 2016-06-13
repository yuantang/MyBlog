package controllers

import (
	"MyBlog/models/class"
	"fmt"
	"strconv"
	"strings"
	//"github.com/russross/blackfriday"
	//"github.com/microcosm-cc/bluemonday"

)

type ArticleController struct {
	BaseController
	ret RET
}

func (c *ArticleController) PageNew() {
	c.CheckLogin()
	c.TplName = "article/new.html"
}
func (c *ArticleController) New() {
	c.CheckLogin()
	u := c.GetSession("user").(class.User)
	a := &class.Article{
		Title:   c.GetString("title"),
		Content: c.GetString("content"),
		Author:  &u,
	}
	n, error := a.Create()
	if error == nil {
		c.ret.OK = true
		c.ret.Content = n
		c.Data["json"] = c.ret
		c.ServeJSON()
		return
	}
	c.ret.Content = fmt.Sprint(error)
	c.Data["json"] = c.ret
	c.ServeJSON()

}


func (c *ArticleController) Get() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	a := &class.Article{Id: id}
	a.ReadDB()
	a.Author.ReadDB()

	a.Replys=class.Reply{Article:a}.Gets()//原始数据


	////unsafe := blackfriday.MarkdownBasic([] byte(a.Content))
	////html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	//a.Content=string(unsafe);

	c.Data["article"] = a
	c.Data["replyTree"]=a.GetReplyTree()
	c.TplName = "article/article.html"
}
func (c *ArticleController) PageEdit() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	a := &class.Article{Id: id}
	a.ReadDB()
	a.Author.ReadDB()
	c.Data["article"] = a
	c.TplName = "article/edit.html"
}
func (c *ArticleController) Edit() {
	c.CheckLogin()
	u := c.GetSession("user").(class.User)
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	a := &class.Article{Id: id}
	a.ReadDB()
	//判读文章作者是不是作者本人
	if u.Id!=a.Author.Id {
		c.DoLogout()
	}
	strs :=strings.Split(c.GetString("tag"),",")
	tags :=[]*class.Tag{}
	for _,v:=range strs{
		tags=append(tags,class.Tag{Name:strings.TrimSpace(v)}.GetOrNew())
	}
	a.Title=c.GetString("title")
	a.Content=c.GetString("content")
	a.Tags=tags

	a.Update()

	c.ret.OK=true
	c.Data["json"]=c.ret
	c.ServeJSON()
}

func (c *ArticleController) Del(){

	c.CheckLogin()
	u := c.GetSession("user").(class.User)
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	a := &class.Article{Id: id}
	a.ReadDB()

	if u.Id!=a.Author.Id{
		c.DoLogout()
	}
	replys:=class.Reply{Article:a}.Gets()//原始数据
	for _,reply:=range replys  {
		reply.Defunct=true
		reply.Update()
	}
	a.Defunct = true
	a.Update()
	c.Redirect("/user/"+a.Author.Id,302)
}

func (c *ArticleController )Archive()  {
	errmsg:=""
	a:=class.Article{}
	if len(c.GetString("tag"))>0 {
		tag:=class.Tag{Name:c.GetString("tag")}.Get()
		if tag==nil{
			errmsg+=fmt.Sprintf("Tag[%s] is not exist.\n",c.GetString("tag"))
		}else {
			a.Tags=[]*class.Tag{tag}
		}
	}
	if len(c.GetString("author"))>0{
		author := class.User{Id: c.GetString("author")}.Get()
		if author == nil {
			errmsg += fmt.Sprintf("User[%s] is not exist.\n", c.GetString("author"))
		} else {
			a.Author = author
		}
	}
	if len(errmsg) == 0 {
		rets := a.Gets()
		//for id,val:=range rets{
		//	unsafe := blackfriday.MarkdownBasic([] byte(val.Content))
		//	val.Content=string(unsafe)
		//	rets[id]=val
		//}
		c.Data["articles"] = rets
	}

	c.Data["err"] = errmsg

	c.TplName = "article/archive.html"
}