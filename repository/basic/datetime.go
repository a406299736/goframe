package basic

import (
	"fmt"
	"strings"
	"time"
)

type DateTime struct {
	time.Time
}

func (dt *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	dt.Time, err = time.ParseInLocation(time.DateTime, s, time.Local)
	return
}

func (dt *DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", dt.Time.Format(time.DateTime))), nil
}
