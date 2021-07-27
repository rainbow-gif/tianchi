package main

import "github.com/gin-gonic/gin"


func room(c *gin.Context) {

}
func enter(w http.ResponseWriter, r *http.Request) {

}func roomLeave(w http.ResponseWriter, r *http.Request) {

}
func roomid(w http.ResponseWriter, r *http.Request) {

}
func users(w http.ResponseWriter, r *http.Request) {

}
func roomList(w http.ResponseWriter, r *http.Request) {

}func user(w http.ResponseWriter, r *http.Request) {

}
func userLogin(w http.ResponseWriter, r *http.Request) {

}
func username(w http.ResponseWriter, r *http.Request) {

}func send(w http.ResponseWriter, r *http.Request) {

}
func retrieve(w http.ResponseWriter, r *http.Request) {

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
	http.HandleFunc("/", home)
	r.Run()
// 默认监听本地8080端口，如果需要更改可以使用 r.Run(":9000")
}
