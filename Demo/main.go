package main

import (
	"Zinx/zface"
	"Zinx/znet"
	"fmt"
)

func main() {
	s := znet.NewServer("[zinx V0.3]")
	s.AddRouter(&PingPost{})
	s.Serve()
}

type PingPost struct {
	znet.BaseRouter
}

// PreHandle Test
func (p *PingPost) PreHandle(request zface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}
}

// Handle Test
func (p *PingPost) Handle(request zface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...\n"))
	if err != nil {
		fmt.Println("call back ping error")
	}

}

// PostHandle Test
func (p *PingPost) PostHandle(request zface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}
}
