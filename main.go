package main

import (
	_ "example.com/m/docs" // 千万不要忘了导入把你上一步生成的docs
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	_ "github.com/go-sql-driver/mysql"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	//"xorm.io/xorm"

	//"time"
	"database/sql"
	//"xorm.io/xorm"
)
var sqlDb *sql.DB           //数据库连接db
var sqlResponse SqlResponse //响应client的数据
type creatroom struct {
	Id      int64     `xorm:"pk autoincr" json:"id"` //指定主键并自增
	Name    string    `json:"name"`
	//Thisroomuser map[string]string

}
//定义结构体(xorm支持双向映射)；没有表，会进行创建
type creatuser struct {
	Id      int64     `xorm:"pk autoincr" json:"id"` //指定主键并自增
	Username 	string    `json:"username"`
	FirstName 	string    `json:"firstName"`
	LastName 	string    `json:"lastName"`
	Email 		string    `json:"email"`
	Password 	string    `json:"password"`
	Phone 		string    `json:"phone"`
	Roomid		string
	//StuNum  string    `xorm:"unique" json:"stu_num"`
	//Name    string    `json:"name"`
	//Age     int       `json:"age"`
	//Created time.Time `xorm:"created" json:"created"`
	//Updated time.Time `xorm:"updated" json:"updated"`

}

//应答体
type SqlResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
func init() {
	//1、打开数据库
	//parseTime:时间格式转换(查询结果为时间时，是否自动解析为时间);
	// loc=Local：MySQL的时区设置
	sqlStr := "root:123456@tcp(127.0.0.1:3306)/xorm?charset=utf8&parseTime=true&loc=Local"
	var err error
	sqlDb, err = sql.Open("mysql", sqlStr)
	if err != nil {
		fmt.Println("数据库打开出现了问题：", err)
		return
	}
	//2、 测试与数据库建立的连接（校验连接是否正确）
	err = sqlDb.Ping()
	if err != nil {
		fmt.Println("数据库连接出现了问题：", err)
		return
	}
	file, err := os.Open("room.sql")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	_, err = sqlDb.Exec(string(content))
	if err != nil {
		return
	}

}


type roomDTO struct {
	name string
}
// GetPostListHandler2 创建房间接口
// @Summary 创建房间接口
// @Tags 房子相关接口
// @Accept application/json
// @Produce text/plain
// @Param name body roomDTO true "Add account"
// @Success 200 string response is room id string
// @Router /room [post]
func room(c *gin.Context) {
	// GET请求参数(query string)：/api/v1/posts2?page=1&size=10&order=time
	// 初始化结构体时指定初始参数
	json := make(map[string]string) //注意该结构接受的内容
	err := c.BindJSON(&json)
	if err != nil {
		return 
	}
	roomname := json["name"]

	log.Printf("%v",&json)
	sqlStr := "insert into room(Name) values (?)"
	ret, err := sqlDb.Exec(sqlStr, roomname)
	if err != nil{
		fmt.Printf("insert failed, err:%v\n", err)
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "写入失败"
		sqlResponse.Data = err
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	id, err := ret.LastInsertId()
	c.String(http.StatusOK, strconv.FormatInt(id, 10))


}
func enter(c *gin.Context) {
	roomid := c.Param("roomid")
	if roomid == "1" {
		c.String(http.StatusOK, "Enter the Room")
	}else {

		c.AbortWithStatus(http.StatusBadRequest)
	}
}
func roomLeave(c *gin.Context) {
	roomid := "1";
	if roomid == "1" {
		c.String(http.StatusOK, "Leave the Room")
	}else {

		c.AbortWithStatus(http.StatusBadRequest)
	}
}
func roomid(c *gin.Context) {
	roomid := c.Param("roomid")
	if roomid == "1" {
		c.String(http.StatusOK, "response is room name string")
	}else {

		c.AbortWithStatus(http.StatusBadRequest)
	}
}
func users(c *gin.Context) {
	roomid := c.Param("roomid")
	allUsers := []string{"fwef","hiuiu"}
	if roomid == "1" {
		c.JSON(http.StatusOK, allUsers)
	}else {

		c.AbortWithStatus(http.StatusBadRequest)
	}
}
func roomList(c *gin.Context) {
	json := make(map[string]int) //注意该结构接受的内容
	err := c.BindJSON(&json)
	if err != nil {
		return
	}
	pageIndex := json["pageIndex"]
	pageSize := json["pageSize"]

	log.Printf("%v",&json)
	allUsers := []map[string]string{{"name":"roomname","id":"rommid"}}
	if pageIndex == pageSize {//假装是这样，其实是分页
		c.JSON(http.StatusOK, allUsers)
	}else {

		c.AbortWithStatus(http.StatusBadRequest)
	}
}
func user(c *gin.Context) {
	json := make(map[string]string) //注意该结构接受的内容
	err := c.BindJSON(&json)
	if err != nil {
		return
	}

	username := json["username"]
	firstName := json["firstName"]
	lastName := json["lastName"]
	email := json["email"]
	password := json["password"]
	phone := json["phone"]
	log.Printf("%s",username+firstName+lastName+email+password+phone)
	c.AbortWithStatus(http.StatusOK)
}
func userLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	log.Printf("%s",username+password)

	c.String(http.StatusOK, uuid.New())
}
func username(c *gin.Context) {
	username  := c.Param("username")
	log.Printf("%s",username)
	c.JSON(http.StatusOK, gin.H{
		"firstName": "string",
		"lastName": "string",
		"email": "string",
		"phone": "string",
	})
}
func send(c *gin.Context) {
	json := make(map[string]string) //注意该结构接受的内容
	err := c.BindJSON(&json)
	if err != nil {
		return
	}
	id := json["id"]
	text := json["text"]
	log.Printf("%s",id + text)
}
func retrieve(c *gin.Context) {
	json := make(map[string]int) //注意该结构接受的内容
	err := c.BindJSON(&json)
	if err != nil {
		return
	}
	pageIndex := json["pageIndex"]
	pageSize := json["pageSize"]
	allUsers := []map[string]string{{"name":"roomname","id":"rommid"}}
	if pageIndex == pageSize {//假装是这样，其实是分页
		c.JSON(http.StatusOK, allUsers)
	}else {

		c.AbortWithStatus(http.StatusBadRequest)
	}
}
func main() {

	r := gin.Default()
	r.POST("/room", room)
	r.PUT("/room/:roomid/enter", enter)
	r.PUT("/roomLeave", roomLeave)
	r.GET("/room/:roomid", roomid)
	r.GET("/room/:roomid/users", users)
	r.POST("/roomList", roomList)
	r.POST("/user", user)
	r.GET("/userLogin", userLogin)
	r.GET("/user/:sername", username)
	r.POST("/message/send", send)
	r.POST("/message/retrieve", retrieve)
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	//http.HandleFunc("/", home)
	r.Run()
	// 默认监听本地8080端口，如果需要更改可以使用 r.Run(":9000")
}
