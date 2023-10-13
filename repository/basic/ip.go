package basic

import "github.com/a406299736/goframe/pkg/tools"

type IpInfo struct {
	Ip string `json:"ip" form:"ip"`
}

func (i *IpInfo) Ip2Long() int64 {
	long, err := tools.IPString2Long(i.Ip)
	if err != nil {
		return 0
	}
	return int64(long)
}
