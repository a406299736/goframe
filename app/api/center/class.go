package center

import (
	"encoding/json"
	"strconv"

	"gitlab.weimiaocaishang.com/weimiao/go-basic/app/pkg/code"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/configs"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/errors"
	"gitlab.weimiaocaishang.com/weimiao/go-basic/pkg/httpclient"
)

const successCode = 20000

// 班级信息
func ClassInfo(raw json.RawMessage) (any interface{}, error2 errors.Er) {
	res, err := httpclient.PostJSON(configs.Get().Center.ClassUrl+"/api/class/info", raw, httpclient.WithHeader("h-app-id", "98"))
	if err != nil {
		return nil, errors.NewErr(100002, err.Error())
	}

	mp := make(map[string]interface{})
	err = json.Unmarshal(res, &mp)
	if err != nil {
		return nil, errors.NewErr(code.JsonParseError, err.Error())
	}

	if _, ok := mp["code"]; !ok {
		return nil, errors.NewErr(code.MapKeyNotExists, err.Error())
	}

	var msg string
	if translateRespCode(mp["code"], successCode) {
		return mp["data"], nil
	} else {
		msg = mp["msg"].(string)
	}

	return nil, errors.NewErr(code.ThirdRespError, msg)
}

// 三方返回respCode 是否等于 三方定义的成功code
func translateRespCode(respCode interface{}, successCode int) bool {
	var i int
	switch respCode.(type) {
	case float64:
		j := respCode.(float64)
		i = int(j)
	case string:
		j := respCode.(string)
		i, _ = strconv.Atoi(j)
	}

	return i == successCode
}
