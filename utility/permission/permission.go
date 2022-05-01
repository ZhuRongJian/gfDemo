package permission

import (
	"strings"
)

//func KeyMatch(key1, key2 string) bool {
//	i := strings.Index(key2, "*")
//	if i == -1 {
//		return key1 == key2
//	}
//	if len(key1) > i {
//		return key1[:i] == key2[:i]
//
//	}
//	return key1 == key2[:i]
//}

// 字符串匹配 请求参数 数据库配置参数
// /user/delete,/user/delete
// /user/delete/8888,/user/delete/*
// /user/delete,"/user/delete/*
func KeyMatch(key1 string, key2 string) bool {
	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}
	if len(key1) > i {
		return key1[:i-1] == key2[:i-1]
	}
	return key1 == key2[:i-1]
}
