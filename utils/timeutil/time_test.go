package timeutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroTimeFormat(t *testing.T) {
	assert.Equal(t, "1970-01-01 08:00:00", TimeFormat(ZeroTime()))
}
