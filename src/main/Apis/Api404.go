package Apis

import (
	"net/http"
	"public.sunibas.cn/go_utils_public/src/utils/AboutServer"
)

func Api404(writer http.ResponseWriter,s string) {
	rObj := AboutServer.ReturnObj{}
	rObj.SetFail("找不到指定接口")
	rObj.Send(writer)
}