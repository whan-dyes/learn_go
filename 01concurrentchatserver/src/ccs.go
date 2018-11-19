package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string //用户发送数据的管道
	Name string      //用户名
	Addr string      //网络地址
}

var onlineMap map[string]Client //保存在线用户

var message = make(chan string)

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C { //给当前客户端发送消息
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ":" + msg
	return
}

func Manager() {
	//给map分配空间
	onlineMap = make(map[string]Client)
	for {
		msg := <-message //没有消息会阻塞
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func HandleConn(conn net.Conn) { //处理用户链接
	defer conn.Close()
	//获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()

	cli := Client{make(chan string), cliAddr, cliAddr}
	//把结构体添加到map
	onlineMap[cliAddr] = cli

	//新开一个协程，专门给当前客户端发送消息
	go WriteMsgToClient(cli, conn)

	//广播某个人在线
	//message<-"["+cli.Addr+"]"+cli.Name+":login"
	message <- MakeMsg(cli, "login hello")

	isQuit := make(chan bool) //对方是否主动退出
	hasData := make(chan bool)
	//新开一个协程，接收用户发送过来的数据
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 { //对断开或对方出问题
				isQuit <- true
				fmt.Println("conn.Read err=", err)
				return
			}
			msg := string(buf[:n-1]) //通过windows nc测试，多一个换行
			if len(msg) == 3 && msg == "who" {
				//遍历map，给当前用户发送所有成员
				conn.Write([]byte("user list:\n"))
				for _, tmp := range onlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write(([]byte(msg)))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				//rename|mike
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename ok\n"))

			} else {
				//转发此内容
				message <- MakeMsg(cli, msg)
			}
			hasData <- true //代表有数据
		}
	}() //别忘了
	for {
		//通过 select 检测channel的流动
		select {
		case <-isQuit:
			delete(onlineMap, cliAddr) //当前用户从map移出
			message <- MakeMsg(cli, "logout")
			return
		case <-hasData:
		case <-time.After(60 * time.Second): //60s后强制下线
			delete(onlineMap, cliAddr) //当前用户从map移出
			message <- MakeMsg(cli, "time out:logout")
			return
		}
	}

}

func main() {
	/*func Listen(net, laddr string) (Listener, error)
	返回在一个本地网络地址laddr上监听的Listener。网络类型参数net必须是面向流的网络：
	"tcp"、"tcp4"、"tcp6"、"unix"或"unixpacket"。参见Dial函数获取laddr的语法。

	type Listener interface {
	    // Addr返回该接口的网络地址
	    Addr() Addr
	    // Accept等待并返回下一个连接到该接口的连接
	    Accept() (c Conn, err error)
	    // Close关闭该接口，并使任何阻塞的Accept操作都会不再阻塞并返回错误。
	    Close() error
	}
	Listener是一个用于面向流的网络协议的公用的网络监听器接口。
	多个线程可能会同时调用一个Listener的方法。
	*/
	listener, err := net.Listen("tcp", ":8800")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listener.Close()

	//新开一个协程，转发消息，只要有消息来了，给map每个成员发送此消息
	go Manager()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accpet err=", err)
			continue
		}
		go HandleConn(conn) //处理用户链接
	}
}
