package statuscode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeToMessage(t *testing.T) {
	msg := CodeToMessage(StatusOK)
	assert.Equal(t, "请求处理成功", msg)

	msg = CodeToMessage(1)
	assert.Equal(t, "Unknown", msg)
}
