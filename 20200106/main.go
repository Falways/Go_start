package main

import (
	"fmt"
	"net"
	"os"
)

func getApplist(conn net.Conn) {
	//接收服务端反馈
	buffer := make([]byte, 1024*10)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), "waiting server back msg error: ", err)
		return
	}
	fmt.Println(conn.RemoteAddr().String(), "receive server back msg: ", string(buffer[:n]))
	// defer conn.Close()
}

func main() {
	server := "127.0.0.1:9123"
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return
	}
	sendInit(conn)
	//getApplist(conn)
}

func sendInit(conn net.Conn)  {
	words := `{"action":"sdp_front_get_applist","data":null}`
	conn.Write([]byte(words))
	fmt.Println("send over")
	// defer conn.Close()
}