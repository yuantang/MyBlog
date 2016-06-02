package class

import (
	"MyBlog/modules"
	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id   int64
	Name string `orm:"index"`
	//反向多对多关系
	Article []*Article `orm:"reverse(many)"`
}

func (t Tag) Get() *Tag {
	o := orm.NewOrm()
	o.QueryTable("tag").Filter("Name", t.Name).One(&t)
	if t.Id == 0 {
		return nil
	}
	return &t
}
func (t Tag) GetOrNew() *Tag {
	o := orm.NewOrm()
	_, _, _ = o.ReadOrCreate(&t, "Name")
	return &t
}

var bscolor []string = []string{"success", "danger", "primary", "warning"}

func (t Tag) RandColor() string {
	return bscolor[modules.RandInt(4)]
}
