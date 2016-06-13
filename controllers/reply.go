package controllers

import (
	"MyBlog/models/class"
	"regexp"
	"strings"
)

type ReplyController struct {
	BaseController
	ret RET
}

func (c *ReplyController) New() {
	c.CheckLogin()
	u := c.GetSession("user").(class.User)

	defer func() {
		c.Data["json"] = c.ret
		c.ServeJSON()
	}()

	article_id, _ := c.GetInt("article_id")

	reply := &class.Reply{
		Author:  &u,
		Article: &class.Article{
			Id: article_id,},

		Content: c.GetString("content"),
	}

	if ok,_:=regexp.MatchString(`^\@\w`,reply.Content);ok {
		reply.ParentId,_=c.GetInt("parent_id")
		reply.Content=strings.SplitN(reply.Content," ",2)[1]
	}
	
	if len(reply.Content) < 1 {
		c.ret.OK = false
		c.ret.Content = "评论不能为空"
		return
	}
	_, err := reply.Create()
	if err != nil {
		c.ret.OK = false
		c.ret.Content = err.Error()
	}
	c.ret.OK = true

}
