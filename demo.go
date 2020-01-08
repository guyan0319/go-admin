package main

import (
	"fmt"
	_ "go-admin/modules/memory" //这里修改成你存放menory.go相应的目录
	"go-admin/modules/session"  //这里修改成你存放session.go相应的目录
	"net/http"
	"unsafe"
)

var globalSessions *session.Manager



func init() {
	var err error
	globalSessions, err = session.NewSessionManager("memory", "goSessionid", 3600)
	if err != nil {
		fmt.Println(err)
		return
	}
	go globalSessions.GC()
	fmt.Println("fd")
}

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("name")
	fmt.Println(cookie)
	if err == nil {
		fmt.Println(cookie.Value)
		fmt.Println(cookie.Domain)
		fmt.Println(cookie.Expires)
	}
	fmt.Fprintf(w, "Hello world!\n") //这个写入到w的是输出到客户端的
}
func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	val := sess.Get("username")
	if val != nil {
		fmt.Println(val)
	} else {
		sess.Set("username", "jerry")
		fmt.Println("set session")
	}
}
func loginOut(w http.ResponseWriter, r *http.Request) {
	//销毁
	globalSessions.SessionDestroy(w, r)
	fmt.Println("session destroy")
}

func main() {
	//http.HandleFunc("/", sayHelloHandler) //	设置访问路由
	//http.HandleFunc("/login", login)
	//http.HandleFunc("/loginout", loginOut) //销毁
	//log.Fatal(http.ListenAndServe(":8010", nil))
	// 设置一个 int64 的数据
	int64_num := int64(6)
	// 将 int64 转化为 int
	int_num := *(*int)(unsafe.Pointer(&int64_num))
	fmt.Println(int_num,int(int64_num))
}