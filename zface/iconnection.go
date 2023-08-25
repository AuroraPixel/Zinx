package zface

import "net"

// IConnection 连接接口
type IConnection interface {
	// Start 启动连接 开始连接工作
	Start()
	// Stop 停止连接 结束当前连接的工作
	Stop()
	// GetTCPConnection 获取当前连接的绑定socket conn
	GetTCPConnection() *net.TCPConn
	// GetConnID 获取当前连接的模块的连接ID
	GetConnID() uint32
	// RemoteAddr 获取客户端的TCP状态
	RemoteAddr() net.Addr
	// Send 发送数据给远程客户端
	Send(data []byte) error
}

// HandleFunc 处理连接的业务方法
type HandleFunc func(*net.TCPConn, []byte, int) error
