package class

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id     string `orm:"pk"`
	Nick   string
	Info   string `orm:null`
	Hobby  string `orm:null`
	Email  string `orm:unique`
	Avator string `orm:null`
	Url    string `orm:null`

	Followers int
	Following int

	Retitme  time.Time `orm:auto_now_add`
	Password string
	Private  int
}

const (
	PR_live = iota
	PR_login
	PR_post
)
const (
	DefaultPvt = 1<<3 - 1
)

//CURD
func  (u User) Get() *User {
	o:=orm.NewOrm()
	err:=o.Read(&u)
	if err==orm.ErrNoRows {
		return nil
	}
	return &u
}
func (u *User) ReadDB() (error error) {
	o := orm.NewOrm()
	error = o.Read(u)
	return
}

func (u *User) Create() (error error) {
	o := orm.NewOrm()
	_, error = o.Insert(u)
	return
}
func (u *User) Update() (error error) {
	o := orm.NewOrm()
	_, error = o.Update(u)
	return
}
func (u *User) Delete() (error error) {
	//	xxx1 & 1110 = xxx0
	//	~x ==> ^x (-1 ^ x)
	u.Private &= ^1
	error = u.Update()
	return
}
func (u *User) ExisId() bool {
	o := orm.NewOrm()
	if error := o.Read(u); error == orm.ErrNoRows {
		return false
	}
	return true
}
func (u User) ExisEmail() bool {
	o := orm.NewOrm()
	return o.QueryTable("user").Filter("Email", u.Email).Exist()
}
