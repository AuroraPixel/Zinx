package znet

import (
	"Zinx/zface"
	"fmt"
	"net"
)

type Connection struct {
	//当前连接的socket TCP
	Conn *net.TCPConn

	//链接的ID
	ConnID uint32

	//当前连接的状态
	IsClosed bool

	//当前连接所绑定的处理业务方法API
	HandleAPI zface.HandleFunc

	//告知当前连接已经退出的/停止 channel
	ExitChan chan bool
}

// NewConnection 初始化连接方法
func NewConnection(conn *net.TCPConn, connID uint32, callBackAPI zface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		IsClosed:  false,
		HandleAPI: callBackAPI,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

// Start 启动连接 开始连接工作
func (c *Connection) Start() {
	fmt.Println("Conn start()...ClientID=", c.ConnID)
	//启动连接读业务操作
	go c.StartReader()
}

// Stop 停止连接 结束当前连接的工作
func (c *Connection) Stop() {
	fmt.Println("Conn stop()...ClientID=", c.ConnID)
	//判断是否已经关闭连接
	if c.IsClosed {
		return
	}
	//关闭连接
	c.IsClosed = true
	c.Conn.Close()
	//关闭管道
	close(c.ExitChan)
}

// GetTCPConnection 获取当前连接的绑定socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取当前连接的模块的连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取客户端的TCP状态
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// Send 发送数据给远程客户端
func (c *Connection) Send(data []byte) error {
	return nil
}

// StartReader 连接读业务操作
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connID=", c.ConnID, "Reader is exit,remote addr is", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端数据到buf周，最大512字节
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}
		//调用当前连接所绑定handleAPI
		if err := c.HandleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID", c.ConnID, "handle is error", err)
			continue
		}
	}
}
