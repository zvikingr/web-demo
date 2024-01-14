package trace

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"web-demo/utils/uuid"
)

const (
	// KeyLocalContext 作为往gin.Context中增加业务上使用trace信息的key
	// 为了尽可能的不和gin本身的key冲突，该值尽量特殊一些
	KeyLocalContext = "___local_context___"
)

// Context 公共请求上下文接口
type Context interface {
	// TraceID 从上下文返回请求链路唯一标识
	TraceID() string

	// String 返回请求上下文中需要输出到log的信息
	String() string
}

// AddTrace add trace context to gin.Context
func AddTrace(c *gin.Context, traceID string) {
	var ctx Context

	if traceID == "" {
		ctx = newContext()
	} else {
		ctx = newContextWithTraceID(traceID)
	}

	c.Set(KeyLocalContext, ctx)
}

// FromContext get local context from gin.Context
// usage:
//
//	log.Errorf("%s||msg=download chart failed||chartName=%s||err=%v", trace.FromContext(c), chartName, err)
func FromContext(c *gin.Context) Context {
	if c == nil {
		return newContext()
	}

	v, ok := c.Get(KeyLocalContext)
	if !ok {
		return newContext()
	}

	ctx, ok := v.(Context)
	if !ok {
		return newContext()
	}

	return ctx
}

// localContext 用于记录请求链路信息，方便日志输出和监控埋点
type localContext struct {
	// 链路追踪唯一标识, 作用范围跨多个服务
	// 通常通过header或服务间自定义方式传入(比如通过协议头传递)
	// 如果上游没有按规则传递，则由当前服务生成
	traceID string

	// // 进入当前服务后生成的唯一ID，作用范围为当前服务
	// // 通常可用来标记单个服务对某一个请求的处理日志
	// spanID string

	// 上下文产生的时间，即结构体的生成时间 或者 请求的到达时间
	startTime time.Time
}

func newContext() Context {
	return &localContext{
		traceID:   uuid.NewUUID(),
		startTime: time.Now(),
	}
}

func newContextWithTraceID(traceID string) Context {
	return &localContext{
		traceID:   traceID,
		startTime: time.Now(),
	}
}

// TraceID return traceID
func (ctx *localContext) TraceID() string {
	return ctx.traceID
}

func (ctx *localContext) getElapseTime() time.Duration {
	return time.Since(ctx.startTime)
}

// String 用于日志输出，注意：这里的输出格式最好和项目整体日志风格保持统一，比如都是用"||"作为分割符号
func (ctx *localContext) String() string {
	return fmt.Sprintf("traceID=%s||elapse=%s", ctx.traceID, ctx.getElapseTime())
}
