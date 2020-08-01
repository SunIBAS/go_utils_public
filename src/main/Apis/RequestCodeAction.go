package Apis

import (
	"net/http"
	"public.sunibas.cn/go_utils_public/src/main/MainTools"
	"public.sunibas.cn/go_utils_public/src/utils/AboutServer"
)

func GetRequestCodeAction(writer http.ResponseWriter,s string,config MainTools.Config) {
	var codes = map[string]int {
		"Success":AboutServer.Success,
		"ERROR":AboutServer.Error,
		"FAIL":AboutServer.Fail,
		"END":AboutServer.End,
	}
	rObj := AboutServer.ReturnObj{}
	rObj.SetContent(codes).SetSuccess("").Send(writer)
}