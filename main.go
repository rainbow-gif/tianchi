package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"log"
	"net/http"
)


func room(c *gin.Context) {
	json := make(map[string]string) //注意该结构接受的内容
	err := c.BindJSON(&json)
	if err != nil {
		return 
	}
	log.Printf("%v",&json)
	c.String(http.StatusOK, "response is room id string")


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
		c.String(http.StatusOK, "Enter the Room")
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
	json := make(map[string]int) //注意该结构接受的内容
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
	c.JSON(http.StatusOK, gin.H{
		"firstName": "string",
		"lastName": "string",
		"email": "string",
		"phone": "string"
	})
}
func send(c *gin.Context) {

}
func retrieve(c *gin.Context) {

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
	//http.HandleFunc("/", home)
	r.Run()
	// 默认监听本地8080端口，如果需要更改可以使用 r.Run(":9000")
}
