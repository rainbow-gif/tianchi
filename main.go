package main

import (
	_ "example.com/m/docs" // 千万不要忘了导入把你上一步生成的docs
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"log"
	"net/http"

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

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
	c.String(http.StatusOK, "response is room id string " + roomname)


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
