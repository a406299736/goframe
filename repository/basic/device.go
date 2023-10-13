package basic

// DeviceInfo 设备信息
type DeviceInfo struct {
	DeviceType string `json:"device_type" form:"device_type"` // 区分大小写
	DeviceId   string `json:"device_id" form:"device_id"`
}

func (d *DeviceInfo) IsIos() bool {
	return d.DeviceType == "ios"
}

func (d *DeviceInfo) IsAndroid() bool {
	return d.DeviceType == "android"
}

func (d *DeviceInfo) IsApplet() bool {
	return d.DeviceType == "applet"
}
