package Apis

import (
	"net/http"
	"public.sunibas.cn/go_utils_public/src/main/MainTools"
	"public.sunibas.cn/go_utils_public/src/utils/AboutServer"
)

func Api404Action(writer http.ResponseWriter,s string,config MainTools.Config) {
	rObj := AboutServer.ReturnObj{}
	rObj.SetFail("找不到指定接口").Send(writer)
}