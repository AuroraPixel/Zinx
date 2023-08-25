package main

import "Zinx/znet"

func main() {
	server := znet.NewServer("[zinx V0.1]")
	server.Serve()
}
