package demohandler

import (
	"net/http"

	"github.com/a406299736/goframe/app/pkg/core"
	"github.com/a406299736/goframe/pkg/errno"
	"github.com/a406299736/goframe/repository/dbrepo/demo"

	"go.uber.org/zap"
)

// demo
func (h *handler) Detail() core.HandlerFunc {
	return func(c core.Context) {
		res, err := h.demoService.Detail(c)

		dm := new(demo.Test1)
		//c.ShouldBindURI(dm) // 获取get参数
		c.ShouldBindJSON(dm) // 获取body json
		c.Info("demo", zap.Any("p11111", dm))

		// 建议统一处理错误, 日志由 app/pkg/core/core.go 统一记录;
		// 非必要情况也可以如下：
		//c.Info("demo1", zap.Any("demo1", "aaaaa1"))
		//c.Logger().Info("demo2", zap.Any("demo2", "bbbbbbb2"))
		if err != nil {
			c.Failed(
				errno.NewError(
					http.StatusOK,
					err.Code(),
					err.Error()),
			)
			return
		}

		c.Success(res)
	}
}

// 方式2 不推荐
func (h *handler1) Detail1() core.HandlerFunc {
	return func(c core.Context) {
		h.demoService.Detail1()
	}
}
