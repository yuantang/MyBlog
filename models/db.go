package models

import (
	"MyBlog/models/class"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.Debug = true


	dbUser := beego.AppConfig.String("DB::user")
	dbPwd  := beego.AppConfig.String("DB::pass")
	dbName := beego.AppConfig.String("DB::name")
	dbAddre:= beego.AppConfig.String("DB::addre")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbUser, dbPwd,dbAddre, dbName,)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbLink)


	//注册数据库驱动（告知使用哪个数据库）
	//orm.RegisterDriver(diverName,DriverType)
	//orm.RegisterDriver("mysql",orm.DRMySQL)
	//注册数据库
	//orm.RegisterDataBase(aliasName,driverName,dataSource,params...)
	//orm.RegisterDataBase("default","mysql","root:123456@tcp(localhost:3306)/myblog?charset=utf8")

	//注册对象模型
	//orm.RegisterModel(models...)
	orm.RegisterModel(new(class.User),new(class.Article),new(class.Tag),new(class.Reply))
	//开启同步
	//orm.RunSyncdb(name string,force bool,verbose bool)
	orm.RunSyncdb("default", false, true)

}
