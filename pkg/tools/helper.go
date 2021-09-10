package tools

import "os"

// 项目绝对路径
func GetProjectAbsolutePath() string {
	s := os.Getenv("PROJECT_PATH")
	if s == "" {
		return "."
	}
	return s
}
