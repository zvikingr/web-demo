package strutil

import (
	"encoding/json"
	"strconv"
)

// ConvertToStr convert object to string
func ConvertToStr(o interface{}) string {
	b, _ := json.Marshal(o)
	return string(b)
}

// StrToInt64 字符串转数字
func StrToInt64(s string) (int64, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return n, nil
}
