package env

import (
	"flag"
	"fmt"
	"strings"
	"testing"
)

var (
	active Environment
	dev    Environment = &environment{value: "dev"}
	test   Environment = &environment{value: "test"}
	rel    Environment = &environment{value: "rel"}
	pro    Environment = &environment{value: "pro"}
)

var _ Environment = (*environment)(nil)

// Environment 环境配置
type Environment interface {
	Value() string
	IsDev() bool
	IsTest() bool
	IsRel() bool
	IsPro() bool
	t()
}

type environment struct {
	value string
}

func (e *environment) Value() string {
	return e.value
}

func (e *environment) IsDev() bool {
	return e.value == "dev"
}

func (e *environment) IsTest() bool {
	return e.value == "test"
}

func (e *environment) IsRel() bool {
	return e.value == "rel"
}

func (e *environment) IsPro() bool {
	return e.value == "pro"
}

func (e *environment) t() {}

func init() {
	env := flag.String("env", "", "请输入运行环境:\n dev:开发环境\n test:测试环境\n rel:预上线环境\n pro:正式环境\n")

	// 防止单元测试时报错
	testing.Init()

	flag.Parse()

	switch strings.ToLower(strings.TrimSpace(*env)) {
	case "dev":
		active = dev
	case "test":
		active = test
	case "rel":
		active = rel
	case "pro":
		active = pro
	default:
		active = test
		fmt.Println("缺少 -env 参数，默认使用test")
	}
}

// Active 当前配置的env
func Active() Environment {
	return active
}
