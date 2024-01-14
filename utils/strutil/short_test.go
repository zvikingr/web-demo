package strutil

import "testing"

func TestShortStr(t *testing.T) {
	str := "http://a.b.c/path/to/resources?a=123&b=333"

	t.Log(ShortStr(str))
}
