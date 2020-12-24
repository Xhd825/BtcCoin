package dbMysql

import (
	"database/sql"
	"fmt"
	"github.com/beego"
	_"github.com/go-sql-driver/mysql"
)
var  Db *sql.DB
func connectDB()  {
	config := beego.AppConfig
	dbDirvername := config.String("db_dirverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbip := config.String("db_ip")
	dbName := config.String("db_name")
	condbUrl := dbUser+":"+dbPassword+"@TCP"+dbip+"/"+dbName+"?charset=utf8"
	DBtype,err := sql.Open(dbDirvername,condbUrl)
	if err != nil {
		panic("数据库连接错误")
	}
	Db=DBtype
	fmt.Println("数据库连接成功")
}