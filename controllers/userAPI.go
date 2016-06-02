package controllers

import (
	"MyBlog/models/class"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	. "fmt"
	"github.com/astaxie/beego/validation"
	"strconv"
	"time"
)

func (c *UserController) API_Profile() {

	type user struct {
		UserId string
		Hobby  []string
	}
	u := user{
		"tom",
		[]string{"code", "chess"},
	}
	c.Data["json"] = u
	c.ServeJSON()
}

type RET struct {
	OK      bool        `json:"success"`
	Content interface{} `json:"content"`
}

func (c *UserController) Register() {

	ret := RET{
		OK:      true,
		Content: "success",
	}
	defer func() {
		c.Data["json"] = ret
		c.ServeJSON()
	}()
	id := c.GetString("userid")
	nick := c.GetString("nick")
	password := c.GetString("password")
	password2 := c.GetString("password2")
	email := c.GetString("email")

	if len(nick) < 1 {
		nick = id
	}
	valid := validation.Validation{}
	valid.Required(id, "UserId")
	valid.Required(password, "Password")
	valid.Required(password2, "Password2")
	valid.Email(email, "Email")
	valid.MaxSize(id, 20, "UserId")
	valid.MaxSize(nick, 30, "Nick")
	switch {
	case valid.HasErrors():
	case password != password2:
		valid.Error("两次密码不一致")
	default:
		u := &class.User{
			Id:       id,
			Nick:     nick,
			Email:    email,
			Password: PwGen(password),
			Retitme:  time.Now(),
			Private:  class.DefaultPvt,
		}
		switch {
		case u.ExisId():
			valid.Error("用户名被占用")
		case u.ExisEmail():
			valid.Error("邮箱被占用")
		default:
			error := u.Create()
			if error == nil {
				return
			}
			valid.Error(Sprintf("%v", error))
		}
	}
	ret.OK = false
	ret.Content = valid.Errors[0].Key + valid.Errors[0].Message

	return
}
func (c *UserController) Login() {
	ret := RET{
		OK:      true,
		Content: "success",
	}
	defer func() {
		c.Data["json"] = ret
		c.ServeJSON()
	}()
	id := c.GetString("userid")
	password := c.GetString("password")
	valid := validation.Validation{}
	//非空判断
	valid.Required(id, "UserId")
	valid.Required(password, "Password")
	valid.MaxSize(id, 20, "userId")
	valid.MaxSize(password,30,"password")

	u := &class.User{Id: id}

	switch {
	case valid.HasErrors():

	case u.ReadDB() != nil:
		valid.Error("用户不存在")
	case PwCheck(password, u.Password) == false:
		valid.Error("密码错误")
	default:
		c.DoLogin(*u)
		ret.OK = true

		return
	}
	ret.Content = valid.Errors[0].Key + valid.Errors[0].Message
	ret.OK = false

	return
}
func (c *UserController) Logout() {
	c.DoLogout()
}
func (c *UserController) Setting() {
	c.CheckLogin()

	switch c.GetString("do") {

	case "info":

		c.SettingInfo()
	case "chpwd":

		c.SettingPwd()
	}
}

func (c *UserController) SettingInfo()  {
	user:=c.GetSession("user").(class.User)
	user.Nick=c.GetString("nick")
	user.Email=c.GetString("email")
	user.Url=c.GetString("url")
	user.Hobby=c.GetString("hobby")
	user.Info=c.GetString("info")
	//数据库更新
	user.Update()
	//session更新
	c.DoLogin(user)
	ret:=RET{
		OK:true,
	}
	c.Data["json"]=ret
	c.ServeJSON()
}
func (c *UserController) SettingPwd()  {
	user:=c.GetSession("user").(class.User)
	//进行密码登录验证
	user.Password=PwGen(c.GetString("pwd2"))
	user.Update()
	c.DoLogin(user)
	ret:=RET{
		OK:true,
	}
	c.Data["json"]=ret
	c.ServeJSON()
}
func PwGen(pass string) string {
	salt := strconv.FormatInt(time.Now().UnixNano()%9000+1000, 10)
	return Base64Encode(Sha1(Md5(pass)+salt) + salt)
}

func PwCheck(pwd, saved string) bool {
	saved = Base64Decode(saved)
	if len(saved) < 4 {
		return false
	}
	salt := saved[len(saved)-4:]
	return Sha1(Md5(pwd)+salt)+salt == saved
}

func Sha1(s string) string {
	return Sprintf("%x", sha1.Sum([]byte(s)))
}

func Md5(s string) string {
	return Sprintf("%x", md5.Sum([]byte(s)))
}

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) string {
	res, _ := base64.StdEncoding.DecodeString(s)
	return string(res)
}
