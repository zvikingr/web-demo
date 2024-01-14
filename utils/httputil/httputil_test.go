package httputil

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"web-demo/utils/statuscode"
)

var (
	r *gin.Engine
)

func init() {
	r = gin.New()

	r.GET("/health", health)
	r.GET("/unHealth", unHealth)
}

func health(c *gin.Context) {
	Response(c, statuscode.StatusOK, nil)
}

func TestResponse(t *testing.T) {
	req, _ := http.NewRequest("GET", "/health", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func unHealth(c *gin.Context) {
	ResponseBadRequest(c, statuscode.AuthError)
}

func TestResponseBadRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "/unHealth", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
