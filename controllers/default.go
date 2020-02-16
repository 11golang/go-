package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"strings"

	//"github.com/astaxie/beego/context"
	"webbee/models"
	"webbee/utils"
)

type MainController struct {
	beego.Controller
}

type LogController struct {
	beego.Controller
}
func (c *LogController) Log(){
	c.TplName = "bottle.html"
	c.Layout = "layout.html"
}
func (this *LogController) Logpost(){
	UserName := this.GetString("username")
	UserPwd := this.GetString("userpwd")
	pwd := models.QueryPwdWithName(UserName)
	//this.Ctx.WriteString(pwd)
	//this.Ctx.WriteString(utils.MD5(UserPwd))
	if strings.Compare(string(pwd), utils.MD5(UserPwd)) >=0{
		this.Data["json"] = map[string]interface{}{
			"code": 1, "message": "登录成功",
		}
	}else{
		this.Data["json"] = map[string]interface{}{
			"code": 0, "message": "登录失败",
		}
	}
	this.ServeJSON()
}
func (c *MainController) Get() {
	c.TplName = "first.html"
	c.Layout = "layout.html"
}

func (c *MainController) Getinfo() {
	c.TplName = "new2.html"
	c.Layout = "layout.html"
}

func (this *MainController) Postinfo() {
	//获取表单信息
	Username := this.GetString("username")
	Userpwd := this.GetString("userpwd")
	Repwd := this.GetString("repwd")
	fmt.Println(Username, Userpwd, Repwd)
	//注册之前先判断该用户名是否已经被注册，如果已经被注册，返回错误
	id := models.QueryUserWithUsername(Username)
	fmt.Println("id", id)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已经存在"}
		this.ServeJSON()

		return
	}
	//注册用户名和密码
	//存储的密码是MD5加密后的数据，那么在登录的验证的时候，也是需要将用户的密码MD5
	//加密之后和存在数据库中的密码进行判断
	Userpwd = utils.MD5(Userpwd)
	user := models.User{Username, Userpwd, Repwd, 0, 0}
	_, err := models.InsertUser(user)
	//this.Ctx.WriteString("插入中...")
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
		panic(err)
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": 1, "message": "注册成功",
		}
		//this.Data["json"]= " OK"
	}
	this.ServeJSON()
}

type GrabdataDisat struct {
	beego.Controller
}

func (this *GrabdataDisat) Getgradbta() {
	//关闭模板渲染
	this.EnableRender = false
	//爬虫入口url
	urls := "https://book.qidian.com/info/1011860223"
	//爬取
	resps := httplib.Get(urls)
	//得到内容
	htmls, err := resps.String()
	if err != nil {
		panic(err)
	}
	this.Ctx.WriteString("小说名称：")
	this.Ctx.WriteString(models.GetnovelName(htmls))
	this.Ctx.WriteString("\n小说作者：")
	this.Ctx.WriteString(models.GetnovelWriter(htmls))
	this.Ctx.WriteString("\n小说简介：")
	this.Ctx.WriteString(models.GetnovelIntroduct(htmls))
	this.Ctx.WriteString("\n小说封面：")
	this.Ctx.WriteString(models.GetnovelPic(htmls))

	maps := make(map[string]string)
	maps["novel_name"] = models.GetnovelName(htmls)
	maps["novel_writer"] = models.GetnovelWriter(htmls)
	maps["novel_introduct"] =models.GetnovelIntroduct(htmls)
	maps["novel_pic"] = models.GetnovelPic(htmls)
	//存入数据库
	models.SqlExcute("novel_info", maps, "")
}
