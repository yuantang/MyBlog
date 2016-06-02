package main

import (
	"github.com/astaxie/beego"
	"encoding/gob"
	"MyBlog/models/class"
	_"MyBlog/models"
	_"MyBlog/routers"
	"strings"

)

func init()  {

	gob.Register(class.User{})
	//注册摸板函数
	beego.AddFuncMap("split",SplitHobby)

}
func main() {

	beego.Run()


}
/*摸板函数*/
func SplitHobby(s string,sep string) []string {
	return strings.Split(s,sep)
}