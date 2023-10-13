package basic

const (
	v10000 = "1.0.0"
)

// VersionInfo 版本信息
type VersionInfo struct {
	Version string `json:"version" form:"version"`
}

func (v *VersionInfo) IsV10000() bool {
	return v.Version == v10000
}
