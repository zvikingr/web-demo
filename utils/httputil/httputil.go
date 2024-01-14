package httputil

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"web-demo/utils/statuscode"
)

// Response 如果业务流程未出错，需要正常返回时，调用该接口
func Response(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": statuscode.CodeToMessage(code),
		"data":    data,
	})
}

// ResponseBadRequest 如果业务流程处理过程中出现错误，调用该接口
// 注意：当前是通过业务code而不是http code来区分客户端错误还是服务端错误
func ResponseBadRequest(c *gin.Context, code int) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    code,
		"message": statuscode.CodeToMessage(code),
		"data":    nil,
	})
}
