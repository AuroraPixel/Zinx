package znet

import "Zinx/zface"

// BaseRouter 实现router先嵌入BaseRouter基类，然后根据这个Router重写
type BaseRouter struct {
}

// PreHandle 处理业务之前的方法
func (r *BaseRouter) PreHandle(request zface.IRequest) {

}

// Handle 处理业务的方法
func (r *BaseRouter) Handle(request zface.IRequest) {

}

// PostHandle 处理业务之后的方法
func (r *BaseRouter) PostHandle(request zface.IRequest) {

}
