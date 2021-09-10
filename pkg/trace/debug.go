package trace

type Debug struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	RTime float64     `json:"r_time"` // 执行时间(单位秒)
}
