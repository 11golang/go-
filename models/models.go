package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"regexp"
	"strings"
)

//匹配小说名称
func GetnovelName(htmls string) string {
	if htmls == "" {
		return ""
	}
	//regexp是正则的包，我们先写好正则的规则
	//<a href="//book.qidian.com/info/1017661402" target="_blank" data-eid="qd_A143" data-bid="1017661402" title="都市大进化时代">都市大进化时代</a>
	reg := regexp.MustCompile(`<em>(.*)</em>`)
	//然后进行匹配 -1 表示全部返回 如果写一个1 他就返回匹配到的第一个  返回是一个[][]string
	result := reg.FindAllStringSubmatch(htmls, -1)
	//如果没有匹配到内容返回空
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}
//<em>心机影帝</em>

//匹配小说作者
func GetnovelWriter(htmls string) string {
	if htmls == "" {
		return ""
	}
	//regexp是正则的包，我们先写好正则的规则
	//<a class="author" href="//me.qidian.com/authorIndex.aspx?id=4406499"
	//data-eid="qd_A144" target="_blank"><img src="//qidian.gtimg.com/qd/images/ico/user.f22d3.png">鸿蒙树</a>
	reg := regexp.MustCompile(`<a href="//my.qidian.com/author/\d*"\s*target="_blank">(.*)</a>`)
	//然后进行匹配 -1 表示全部返回 如果写一个1 他就返回匹配到的第一个  返回是一个[][]string
	result := reg.FindAllStringSubmatch(htmls, -1)
	//如果没有匹配到内容返回空
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}
//匹配小说简介
func GetnovelIntroduct(htmls string) string {
	if htmls == "" {
		return ""
	}
	//<p>灵气复苏之后，世界巨变，人类走向进化时代！</p>
	reg := regexp.MustCompile(`<p class="intro">(.*)</p>`)
	//然后进行匹配 -1 表示全部返回 如果写一个1 他就返回匹配到的第一个  返回是一个[][]string
	result := reg.FindAllStringSubmatch(htmls, -1)
	//如果没有匹配到内容返回空
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}
//<p class="intro">心潮澎湃，无限幻想，迎风挥击千层浪，少年不败热血！</p>
//匹配小说图片
func GetnovelPic(htmls string) string {
	if htmls == "" {
		return ""
	}
	//<img class="lazy" src="//bookcover.yuewen.com/qdbimg/349573/1017661402/90" data-original=
	//	"//bookcover.yuewen.com/qdbimg/349573/1017661402/90" alt="都市大进化时代" style="display: inline;">
	reg := regexp.MustCompile(`<img src="(.*?)">`)
	//然后进行匹配 -1 表示全部返回 如果写一个1 他就返回匹配到的第一个  返回是一个[][]string
	result := reg.FindAllStringSubmatch(htmls, -1)
	//如果没有匹配到内容返回空
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}
//<img src="//bookcover.yuewen.com/qdbimg/349573/1017561454/180
//">
// 数据库插入修改
func SqlExcute(table string, data map[string]string, where string) (num int64, err error) {
	if table == "" {
		err = errors.New("表名不能为空")
		num = 0
		return
	}
	if data == nil {
		err = errors.New("表名不能为空")
		num = 0
		return
	}
	o := orm.NewOrm()
	sql := ""
	if where == "" {
		var field []string
		var value []string
		for k, v := range data {
			field = append(field, "`"+k+"`")
			if v != "" {
				value = append(value, "'"+v+"'")
			} else {
				value = append(value, "null")
			}
		}
		field_str := strings.Join(field, ",")
		value_str := strings.Join(value, ",")
		sql = "insert into " + table + "(" + field_str + ") values(" + value_str + ")"

	} else {
		var set []string
		for k, v := range data {
			if v != "" {
				set = append(set, "`"+k+"`='"+v+"'")
			} else {
				set = append(set, "`"+k+"`=null")
			}
		}
		set_str := strings.Join(set, ",")
		sql = "update " + table + " set " + set_str + " where " + where

	}
	res, err2 := o.Raw(sql).Exec()
	if err2 != nil {
		err = err2
		num = 0
	} else {
		num, err = res.LastInsertId()
	}
	return
}
