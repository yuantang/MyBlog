package class

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Id      int
	Title   string `orm:"size(80)"`
	Content string `orm:type(text)`
	Author  *User  `orm:"rel(fk)"` //外键

	NumReplys int     //文章评论数
	Replys []*Reply `orm:"-"`  //文章评论
	NumViews  int

	Tags    []*Tag    `orm:"rel(m2m)"`  //正向关系多对多
	Time    time.Time `orm:"auto_now_add;type(datetime)"`
	Defunct bool
}

func (a *Article) ReadDB() (error error) {
	o := orm.NewOrm()
	error = o.Read(a)
	if error != nil {
		beego.Info(error)
	}
	_,_=o.LoadRelated(a,"tags")
	return
}
func (a *Article) Create() (n int64, error error) {
	o := orm.NewOrm()
	n, error = o.Insert(a)
	if error != nil {
		beego.Info(error)
	}
	return
}
func (a Article) Update() (error error) {
	o := orm.NewOrm()
	_, error = o.Update(&a)
	if error != nil {
		beego.Info(error)
	}
	m2m:=o.QueryM2M(&a,"Tags")
	//m2m.Clear()
	//m2m.Add(a.Tags)
	old:=Article{Id:a.Id}
	_,_=o.LoadRelated(&old,"Tags")

	//inster
	VI:
	for _,vi:=range a.Tags {
		for _,vj:=range old.Tags {
			if vi.Id == vj.Id {
				continue VI
			}

		}
		m2m.Add(vi)
	}
	//del
	VD:
	for _,vi:=range old.Tags{
		for _,vj:=range a.Tags {
			if vi.Id==vj.Id {
				continue VD
			}
		}
		m2m.Remove(vi)
	}


	return
}
func (a Article) Delete() (error error) {
	a.Defunct = true
	error = a.Update()
	return
}
//获取属性相同的所有文章
func (a Article) Gets() (rets []Article) {
	o := orm.NewOrm()
	//o.QueryTable("article").Filter("Author", a.Author).Filter("defunct", 0).All(&ret)
	qs:=o.QueryTable("article")
	if a.Author!=nil {
		qs=qs.Filter("Author",a.Author)
	}
	if len(a.Tags)==1 {
		qs=qs.Filter("Tags__Tag",a.Tags[0])
	}
	qs=qs.Filter("defunct",0)
	//Author
	qs=qs.RelatedSel()

	qs.All(&rets)
	//Tag
	for i:=range rets{
		_,_=o.LoadRelated(&rets[i],"Tags")
	}

	return
}
//获取评论树
func (a *Article) GetReplyTree() (rets []*ReplyTree) {
	replys:=Reply{Article:a}.Gets()
	m:=make(map[int]*ReplyTree)
	for _,reply:=range replys{
		tr:=&ReplyTree{
			Reply:reply,
			Childs:make([]*ReplyTree,0),
		}
		m[tr.Id]=tr
		//如果parentId是0 就是直接对文章的评论
		if reply.ParentId==0 {
			rets=append(rets,tr)
		}else{
			m[reply.ParentId].Childs=append(m[reply.ParentId].Childs,tr)
		}
	}
	return
}
