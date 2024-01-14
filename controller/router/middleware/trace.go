package middleware

import (
	"github.com/gin-gonic/gin"

	"web-demo/utils/log"
	"web-demo/utils/trace"
)

// KeyTraceID 用来在header中传递traceID的key
const KeyTraceID = "X-TraceID"

// Tracer trace middleware
func Tracer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			requestUrl = c.Request.URL
			httpMethod = c.Request.Method
		)

		traceID := c.GetHeader(KeyTraceID)
		trace.AddTrace(c, traceID)

		ctx := trace.FromContext(c)
		log.Infof("%s||request_in||method=%s||url=%s||%s=%s", ctx, httpMethod, requestUrl, KeyTraceID, traceID)

		c.Next()

		c.Header(KeyTraceID, ctx.TraceID())

		log.Infof("%s||request_out||method=%s||url=%s", ctx, httpMethod, requestUrl)
	}
}
