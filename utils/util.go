package utils

import (
	"crypto/md5"
	"github.com/astaxie/beego/orm"
	_ "github.com/Go-SQL-Driver/MySQL"
	//"database/sql/driver"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"log"
)

var db *sql.DB

func InitMysql() {
	fmt.Println("InitMysql...")
	driverName := beego.AppConfig.String("driverName")
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1
		CreatTableWithUser()
	}
}

//操作数据库,这是一个通用的方法，Exec表示执行的意思
func ModifyDB(sql string, args ...interface{}) (int64, error) { //sql表示要执行的mysql语句
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		return 0, err
	}
	return count, nil
}

//创建用户表
func CreatTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users (
  username varchar(50) DEFAULT NULL,
  userpwd varchar(150) DEFAULT NULL,
  repwd varchar(150) DEFAULT NULL,
  status int(11) NOT NULL DEFAULT '0',
  id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY 
) ;`
	ModifyDB(sql)
}

//利用原本QueryRow方法查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

//MD5语法，对密码进行加密
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str

}

type Info struct {
	Id                   int64  `orm:"pk" from:"id"`
	Novel_name string
	Novel_writer string
	Novel_introduct string
	Novel_pic string
}
func init() {
	//注册mysql驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterModel(new(Info))
	//连接数据库
	orm.RegisterDataBase("default", "mysql", "root:root@(localhost:3306)/test?charset=utf8")
	//生成数据库表表单映射
	orm.RegisterModelWithPrefix("novel_", new(Info))
	//开启自动建表
	orm.RunSyncdb("default", false, false)
	//开启ormdebug模式
	orm.Debug = true
}
