package tools

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"html"
	"math/rand"
	"regexp"
	"strings"
	"unicode"
)

func FmtPrintf(format string, val ...any) {
	fmt.Printf(carbon.Now().ToDateTimeString()+" "+format+" \n", val...)
}

// Md5 进行md5加密
func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))

	return hex.EncodeToString(m.Sum(nil))
}

// Trims 去除字符串左右字符
// s 目标字符串
// ts 要去除的特殊字符串
// return 去除后的字符串
func Trims(s string, ts []string) string {
	for i := 0; i < len(ts); i++ {
		for _, v := range ts {
			s = strings.Trim(s, v)
		}
	}
	return s
}

// TrimHtml 去除html标签
func TrimHtml(src string) string {
	src = html.UnescapeString(src)
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	//src = re.ReplaceAllString(src, "\n")
	src = re.ReplaceAllString(src, "")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

// SubStr 截取字符串
func SubStr(str string, start, length int) string {
	r := []rune(str)

	maxIndex := len(r) - 1
	if maxIndex < start {
		return ""
	}
	if maxIndex >= start && maxIndex < start+length {
		return string(r[start:])
	}

	return string(r[start : start+length])
}

// UcFirst 将字符串首字母大写
func UcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	firstRune := []rune(s)[0]
	if unicode.IsLower(firstRune) {
		firstRune = unicode.ToUpper(firstRune)
	}
	return string(firstRune) + s[1:]
}

// StrRev 字符串翻转
func StrRev(str string) string {
	runes := []rune(str)
	runesLen := len(runes)

	for i := 0; i < runesLen/2; i++ {
		runes[i], runes[runesLen-1-i] = runes[runesLen-1-i], runes[i]
	}

	return string(runes)
}

// StrShuffle 打乱字符串
func StrShuffle(str string) string {
	runes := []rune(str)
	runesLen := len(runes)

	for i := runesLen - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
