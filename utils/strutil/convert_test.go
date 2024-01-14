package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type user struct {
	Name string `json:"name"`
}

func TestCovertToStr(t *testing.T) {
	u := &user{
		Name: "admin",
	}

	str := ConvertToStr(u)
	assert.Equal(t, "{\"name\":\"admin\"}", str)
}
