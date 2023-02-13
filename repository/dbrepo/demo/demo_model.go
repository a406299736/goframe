package demo

import (
	"encoding/json"
	"time"
)

// 示例
type Demo struct {
	Id          int32     `json:"id"`                    // 主键
	Username    string    `json:"username"`              // 用户名
	Password    string    `json:"password"`              // 密码
	Nickname    string    `json:"nickname"`              // 昵称
	Mobile      string    `json:"mobile"`                // 手机号
	IsUsed      int32     `json:"isUsed"`                // 是否启用 1:是  -1:否
	IsDeleted   int32     `json:"isDeleted"`             // 是否删除 1:是  -1:否
	CreatedAt   time.Time `json:"createdAt" gorm:"time"` // 创建时间
	CreatedUser string    `json:"createdUser"`           // 创建人
	UpdatedAt   time.Time `json:"updatedAt" gorm:"time"` // 更新时间
	UpdatedUser string    `json:"updatedUser"`           // 更新人
}

type Test1 struct {
	Id          int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Aspirations string    `gorm:"column:aspirations;NOT NULL" json:"aspirations"`                   
	Map         string    `gorm:"column:map;NOT NULL" json:"map"`                                   // 网站地图
	Call        string    `gorm:"column:call;NOT NULL" json:"call"`                                 // 联系我们
	Recruit     string    `gorm:"column:recruit;NOT NULL" json:"recruit"`                           // 人才招聘
	Updated     time.Time `gorm:"column:updated;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated"` // 更新时间
	Created     time.Time `gorm:"column:created;default:CURRENT_TIMESTAMP;NOT NULL" json:"created"` // 创建时间
}

func (m *Test1) TableName() string {
	return "test1"
}

func (m *Test1) String() string {
	bytes, _ := json.Marshal(m)
	return string(bytes)
}
