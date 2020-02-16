package models

import (
	"fmt"
	"webbee/utils"
)

type User struct {
	Username string `json: username`
	Userpwd  string `json:userpwd`
	Repwd    string `json:repwd`
	Status   int    `json:status`
	Id       int    `json:id`
}

//数据库操作
//插入
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username,userpwd,repwd) values(?,?,?)",
		user.Username, user.Userpwd, user.Repwd)
}

//按条件查询(，主要在users后面条件的设立
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//根据用户名查询id
//法一独自实现该功能
//func QueryUserWithUsername(username string) int {
//	sql :=fmt.Sprintf("select id from users where username ='%s' ",username)
//	row :=utils.QueryRowDB(sql)
//	id :=0
//	row.Scan(&id)
//	return id
//}
//方法二通过调用通用函数，条件查询函数来实现该功能
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username= '%s'", username)
	return QueryUserWightCon(sql)
}

//根据用户名和密码，查询id
func QueryUserWithParam(username, userpwd string) int {
	sql := fmt.Sprintf(" where useername ='%s'&& userpwd ='%s'", username, userpwd)
	return QueryUserWightCon(sql)
}
func QueryPwdWithName(name string) string {
	sql := fmt.Sprintf("select userpwd from users where username ='%s'", name)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	pwd :=""
	row.Scan(&pwd)
	return pwd
}
