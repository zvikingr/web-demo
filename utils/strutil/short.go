package strutil

import (
	"github.com/spaolacci/murmur3"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func encode(number int) string {
	if number == 0 {
		return string(alphabet[0])
	}

	chars := make([]byte, 0)

	length := len(alphabet)

	for number > 0 {
		result := number / length
		remainder := number % length
		chars = append(chars, alphabet[remainder])
		number = result
	}

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

// ShortStr Convert a long string to a shorter string
// eg:
//
//	input: http://a.b.c/path/to/resources?a=123&b=333
//	outPut: tMyog
func ShortStr(v string) string {
	incr := murmur3.Sum32([]byte(v))
	return encode(int(incr))
}
