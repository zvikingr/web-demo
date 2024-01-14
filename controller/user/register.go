package user

import (
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) error {
	handle := NewHandle()

	r.PUT("/users", handle.UpdateUser)

	return nil
}
