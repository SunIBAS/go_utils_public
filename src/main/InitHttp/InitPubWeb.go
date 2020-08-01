package InitHttp

import (
	"encoding/json"
	"net/http"
	"public.sunibas.cn/go_utils_public/src/main/Sqls"
	"public.sunibas.cn/go_utils_public/src/utils/AboutServer"
	"public.sunibas.cn/go_utils_public/src/utils/DirAndFile"
)

// 这里的目的是发布一个本地的网站，访问的根部是 pubWeb

var (
	pubWebRoot = "/pubWeb"
	pubWebMap = map[string] string{}
)

func InitPubWeb() {
	http.HandleFunc(pubWebRoot + "Add", func(writer http.ResponseWriter, request *http.Request) {
		params,err := AboutServer.PostParams(request)
		rObj := AboutServer.ReturnObj{ }
		if err == nil {
			if "add" == params.Method {
				var pubweb Sqls.PubWebEntity
				json.Unmarshal([]byte(params.Content),&pubweb)
				if t,err := DirAndFile.CheckTypeAndExist(pubweb.Dir);err == nil && t == 1 {
					http.Handle(pubWebRoot + "/" + pubweb.PerPath,
						http.StripPrefix(pubWebRoot + "/" + pubweb.PerPath,
							http.FileServer(http.Dir(pubweb.Dir))))
					pubWebMap[pubWebRoot + "/" + pubweb.PerPath] = pubweb.Dir
					rObj.SetSuccess("定义成功,链接为 :" + pubWebRoot + "/" + pubweb.PerPath).Send(writer)
				} else {
					if err != nil {
						config.Logger.Printf("[pubweb][fail][add] 解析错误:%s",err.Error())
						rObj.SetFail("[pubweb][fail][add] 解析错误:" + err.Error()).Send(writer)
					} else {
						tis := "不存在"
						if t != 0 {
							tis = "文件"
						}
						config.Logger.Printf("[public][fail][add] 路径类型不是文件夹而是:%s",tis)
						rObj.SetFail("[public][fail][add] 路径类型不是文件夹而是:" + tis).Send(writer)
					}
				}
			}
		}
	})
}
