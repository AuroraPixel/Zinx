package znet

import (
	"Zinx/zface"
	"fmt"
	"net"
)

// Server iServer的接口实现
type Server struct {
	// 服务器名称
	Name string
	// 服务器绑定的IP版本
	IPVersion string
	// 服务器监听的IP
	IP string
	// 服务器监听的端口
	Port int
	// Router成员
	Router zface.IRouter
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP :%s, Port %d,is starting\n", s.IP, s.Port)
	//1.获取TCP的Addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error:", err)
		return
	}
	//2.监听服务器的地址
	listenner, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("listen", s.IPVersion, "err", err)
		return
	}
	fmt.Println("start Zinx server succ", s.Name, "succ Listening ...")
	//3.阻塞等待客户端连接，处理客户端连接业务（读写）
	var cid uint32
	cid = 0
	for {
		//如果有客户端连接过来，阻塞返回
		conn, err := listenner.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err", err)
			continue
		}
		//已经与客户端建立和conn进行绑定得到我们链接模块
		dealConn := NewConnection(conn, cid, s.Router)
		cid++

		//启动当前的链接业务处理
		go dealConn.Start()

	}

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()

	//TODO 做一些启动服务之后的额外服务
	//阻塞
	select {}
}

// AddRouter 添加路由成功
func (s *Server) AddRouter(router zface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Succ!")
}

// NewServer 初始化服务器
func NewServer(name string) zface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
	return s
}
