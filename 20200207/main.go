package main

import (
"log"
"net/http"
)

func main() {
	//注册一个函数，响应某一个路由
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type","application/json")
		w.Write([]byte("{'username':'xuhang'}"))
	})
	//这里可以单独写一个函数传递给当前的路由
	http.HandleFunc("/bye", SayBye)
	log.Println("Start version v1")
	log.Fatal(http.ListenAndServe(":4000", nil))

}
func SayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye bye, this is version v1"))
	//进行一个流式传递，将字符串转换为byte类型
}
