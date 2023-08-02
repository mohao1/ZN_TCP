package znet

import "ZN/ZN/ziface"

// Router 实现IRouter接口
// 被开发者继承
// 因为有些业务不需要用到PreHandle()或者PostHandle避免使用时候把方法全部的重写
type Router struct {
}

// PreHandle 处理conn业务之前的方法
func (r *Router) PreHandle(request ziface.IRequest) {}

// Handle 处理conn业务的方法
func (r *Router) Handle(request ziface.IRequest) {}

// PostHandle 处理conn业务之后的方法
func (r *Router) PostHandle(request ziface.IRequest) {}
