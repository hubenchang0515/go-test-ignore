package main

import (
	"fmt"
	"regexp"
	"strings"
)

/* Note:
 * +build FLAG1 FLAG2              : FLAG1 or FLAG2
 * +build FLAG1,FLAG2              : FLAG1 and FLAG2
 * +build FLAG1,FLAG2 FLAG3,FLAG4  : (FLAG1 and FLAG2) or (FLAG3 and FLAG4)
 */

// 清除逗号两侧的空白字符
func CommaClean(str string) string {
	reg, _ := regexp.Compile(`\s*,\s*`)
	return reg.ReplaceAllString(str, ",")
}

// 判断一个单独的FLAG是否是指定的值
func CheckSignleBuildFlag(flag string, str string) bool {
	str = strings.TrimSpace(str)
	exp := fmt.Sprintf(`^//\s*\+build\s+%s$`, flag)
	matched, _ := regexp.MatchString(exp, str)
	return matched
}

// 判断一个组合FLAGS中是否以AND的方式包含指定的值
func CheckMultiBuildFlag(flag string, str string) bool {
	str = strings.TrimSpace(str)
	str = CommaClean(str)
	exp := fmt.Sprintf(`^\s*//\s*\+build(\s+|\s+[^ ]*,)%s(\s*|,[^ ]*)$`, flag)
	matched, _ := regexp.MatchString(exp, str)
	return matched
}

// 在组合FLAGS中去除指定的值
func RemoveFlag(flag string, str string) string {
	if CheckSignleBuildFlag(flag, str) {
		return ""
	} else if CheckMultiBuildFlag(flag, str) {
		exp := fmt.Sprintf(`%s,{0,1}\s*`, flag)
		reg, _ := regexp.Compile(exp)
		return reg.ReplaceAllString(str, "")
	} else {
		return str
	}
}
