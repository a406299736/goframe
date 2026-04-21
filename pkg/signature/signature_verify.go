package signature

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/golang-module/carbon/v2"
)

func (s *signature) Verify(authorization, date string, path string, method string, params url.Values) (ok bool, err error) {
	if date == "" {
		err = fmt.Errorf("date required")
		return
	}

	if path == "" {
		err = fmt.Errorf("path required")
		return
	}

	if method == "" {
		err = fmt.Errorf("method required")
		return
	}

	methodName := strings.ToUpper(method)
	if !methods[methodName] {
		err = fmt.Errorf("method param error")
		return
	}

	if carbon.Parse(date).DiffInSeconds(carbon.Now()) > int64(s.ttl/time.Second) {
		err = fmt.Errorf("date exceeds limit %v", s.ttl)
		return
	}

	// Encode() 方法中自带 sorted by key
	sortParamsEncode, err := url.QueryUnescape(params.Encode())
	if err != nil {
		err = fmt.Errorf("url QueryUnescape error %v", err)
		return
	}

	buffer := bytes.NewBuffer(nil)
	buffer.WriteString(path)
	buffer.WriteString(delimiter)
	buffer.WriteString(methodName)
	buffer.WriteString(delimiter)
	buffer.WriteString(sortParamsEncode)
	buffer.WriteString(delimiter)
	buffer.WriteString(date)

	// 对数据进行 hmac 加密，并进行 base64 encode
	hash := hmac.New(sha256.New, []byte(s.secret))
	hash.Write(buffer.Bytes())
	digest := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	ok = authorization == fmt.Sprintf("%s %s", s.key, digest)
	return
}
