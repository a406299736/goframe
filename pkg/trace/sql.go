package trace

type SQL struct {
	Timestamp string  `json:"timestamp"`     // 时间，格式：2006-01-02 15:04:05
	Stack     string  `json:"stack"`         // 文件地址和行号
	SQL       string  `json:"sql"`           // SQL 语句
	Rows      int64   `json:"rows_affected"` // 影响行数
	RTime     float64 `json:"r_time"`        // 执行时长(单位秒)
}
