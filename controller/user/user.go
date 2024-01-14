package user

import (
	"github.com/gin-gonic/gin"

	"web-demo/logic"
	"web-demo/utils/httputil"
	"web-demo/utils/log"
	"web-demo/utils/statuscode"
	"web-demo/utils/strutil"
	"web-demo/utils/trace"
)

func (h Handle) UpdateUser(c *gin.Context) {
	ctx := trace.FromContext(c)

	req := new(logic.UserRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		log.Warnf("%s||decode request failed||req=%s||err=%v", ctx, strutil.ConvertToStr(req), err)
		httputil.ResponseBadRequest(c, statuscode.ParamsInvalid)
		return
	}

	err := logic.UpdateUser(ctx, req.UserName, req.Password)
	if err != nil {
		httputil.ResponseBadRequest(c, statuscode.ServerInternalError)
		return
	}

	httputil.Response(c, statuscode.StatusOK, nil)
}
