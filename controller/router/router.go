package router

import (
	"net/http"

	"web-demo/controller/router/middleware"
	"web-demo/controller/user"

	"github.com/gin-gonic/gin"
)

// InitRouter init web engine middleware and load routers
func InitRouter(env string) (http.Handler, error) {
	r := gin.New()
	if env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 加载全局中间件
	r.Use(middleware.Tracer())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	return r, LoadRouterRules(r)
}

// LoadRouterRules load all service route from every group
func LoadRouterRules(r *gin.Engine) error {
	if err := user.Register(r); err != nil {
		return err
	}

	return nil
}
