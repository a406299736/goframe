package com

const (
	v10000 = "1.00.00"
)

// VersionInfo 版本信息
type VersionInfo struct {
	Version string `json:"version" form:"version"`
}

func (v *VersionInfo) IsV10000() bool {
	return v.Version == v10000
}
