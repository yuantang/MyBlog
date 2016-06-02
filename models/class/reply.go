package class

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)
type Reply struct  {

	Id int
	Content string `orm:"type(text)"`
	Article *Article `orm:"rel(fk)"`//文章是外键
	Author *User `orm:"rel(fk)"`
	ParentId int
	Time time.Time `orm:"auto_now_add"`//创建评论时自动添加时间
	Defunct bool
}

type ReplyTree struct  {
	*Reply
	Childs [] *ReplyTree

}


func (r *Reply)Create()  (n int64,err error) {
	o:=orm.NewOrm()
	if n,err=o.Insert(r);err!=nil {
		beego.Info(err)
	}
return
}
func (r Reply)Gets() (rets []*Reply)  {

	o:=orm.NewOrm()
	qs:=o.QueryTable("reply")
	if r.Article!=nil {
		qs=qs.Filter("article_id",r.Article.Id)
	}
	if r.Author!=nil {
		qs=qs.Filter("author_id",r.Author.Id)
	}
	qs=qs.Filter("defunct",0)
	qs.All(&rets)
	for k:=range rets{
		rets[k].Author.ReadDB()
		rets[k].Article.ReadDB()
	}

	return
}