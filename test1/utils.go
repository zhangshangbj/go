package test1

import (
	"crypto/md5"
	"fmt"
)

// Md5 返回一个字符串的md5值 类型为string
func Md5(text string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}
