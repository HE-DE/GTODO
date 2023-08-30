package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5加密小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	temp := h.Sum(nil)
	return hex.EncodeToString(temp)
}

// Md5加密大写
func Md5EncodeUpper(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 加密
func MakePassword(plainpwd string) string {
	return Md5Encode(plainpwd)
}

// 解密
func VaildPassword(plainpwd, password string) bool {
	return Md5Encode(plainpwd) == password
}
