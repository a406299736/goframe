package test1

import "time"

// Test1
//
//go:generate gormgen -structs Test1-input .
type Test1 struct {
	Id          int32     //
	Aspirations string    // 微淼心声
	Map         string    // 网站地图
	Call        string    // 联系我们
	Recruit     string    // 人才招聘
	Updated     time.Time `gorm:"time"` // 更新时间
	Created     time.Time `gorm:"time"` // 创建时间
}
