package tools

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/golang-module/carbon"
	"log"
	"os"
	"strings"
)

func FmtPrintf(format string, val ...any) {
	fmt.Printf(carbon.Now().ToDateTimeString()+" "+format+" \n", val...)
}

// 文件中追加内容
func FilePutContents(filename string, data string) (n int) {
	fileHandle, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("open file error :", err)
		return
	}
	defer fileHandle.Close()
	buf := bufio.NewWriter(fileHandle)
	n, _ = buf.WriteString(data)
	err = buf.Flush()
	if err != nil {
		log.Println("flush error :", err)
	}

	return
}

// 文件是否存在
func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// 进行md5加密
func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))

	return hex.EncodeToString(m.Sum(nil))
}

// 去除字符串左右字符
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
