package zface

type IRouter interface {
	//PreHandle 处理业务之前的方法
	PreHandle(request IRequest)
	//Handle 处理业务的方法
	Handle(request IRequest)
	//PostHandle 处理业务之后的方法
	PostHandle(request IRequest)
}
