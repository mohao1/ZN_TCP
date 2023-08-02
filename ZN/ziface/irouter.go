package ziface

type IRouter interface {
	// PreHandle 处理conn业务之前的方法
	PreHandle(request IRequest)
	// Handle 处理conn业务的方法
	Handle(request IRequest)
	// PostHandle 处理conn业务之后的方法
	PostHandle(request IRequest)
}
