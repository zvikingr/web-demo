package trace

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const testTraceID = "123"

type testContext struct{}

func TestAddTrace(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	AddTrace(c, "")
	_, ok := c.Get(KeyLocalContext)
	assert.Equal(t, true, ok)

	AddTrace(c, testTraceID)
	v, ok := c.Get(KeyLocalContext)
	assert.Equal(t, true, ok)

	ctx, ok := v.(Context)
	assert.Equal(t, true, ok)

	assert.Equal(t, testTraceID, ctx.TraceID())
}

func TestFromContext(t *testing.T) {
	FromContext(nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	FromContext(c)

	AddTrace(c, testTraceID)
	ctx := FromContext(c)
	assert.Equal(t, testTraceID, ctx.TraceID())
	t.Log(ctx.String())

	c.Set(KeyLocalContext, &testContext{})
	FromContext(c)
	_, ok := c.Get(KeyLocalContext)
	assert.Equal(t, true, ok)
}
