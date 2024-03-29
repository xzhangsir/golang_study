package ziface

import "net"

//定义连接接⼝
type IConnection interface {
	//启动连接，让当前连接开始⼯工作
	Start()
	//停⽌连接，结束当前连接状态
	Stop()
	//获取当前连接ID
	GetConnID() uint32
	//从当前连接获取原始的socket TCPConn
	GetTCPConnection() *net.TCPConn
	//获取远程客户端地址信息
	RemoteAddr() net.Addr
	//直接将Message发送数据给远程的TCP客户端
	SendMsg(msgId uint32, data []byte) error
}

//定义一个统⼀处理链接业务的接⼝
//第⼀个参数是socket原⽣链接，第⼆个参数是客户端请求的数据，第三个参数是客户端请求的数据长度
type HandFunc func(*net.TCPConn, []byte, int) error
