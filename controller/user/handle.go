package user

// Handle 实现当前分组版本下需要的所有接口处理句柄
// 如果有各接口函数通用的配置，可以定义成handle的成员变量
type Handle struct{}

// NewHandle create handle
func NewHandle() *Handle {
	return new(Handle)
}
