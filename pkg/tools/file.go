package tools

import (
	"bufio"
	"log"
	"os"
)

// FilePutContents 文件中追加内容
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

// CheckFileExist 文件是否存在
func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
