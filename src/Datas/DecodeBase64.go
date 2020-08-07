package Datas

import (
	"encoding/base64"
	"net/url"
	"strings"
)

// 可以通过修改底层url.QueryEscape代码获得更高的效率，很简单
func encodeURIComponent(str string) string {
	r := url.QueryEscape(str)
	r = strings.Replace(r, "+", "%20", -1)
	return r
}

func DecodeBase64(content string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return ""
	} else {
		return string(decodeBytes)
	}
}
func DecodeBase64AndURI(content string) string {
	ret,err := url.QueryUnescape(strings.Replace(DecodeBase64(content), "+", "%2b",-1))
	if err != nil {
		return ""
	}
	return ret
}

func EncodeBase64(content string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(content)))
}
func EncodeBase64AndURI(content string) string {
	return EncodeBase64(url.QueryEscape(content))
}