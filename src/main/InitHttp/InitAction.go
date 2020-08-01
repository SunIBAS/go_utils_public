package InitHttp

import (
	"net/http"
	"public.sunibas.cn/go_utils_public/src/main/Apis"
	"public.sunibas.cn/go_utils_public/src/main/MainTools"
	"public.sunibas.cn/go_utils_public/src/utils/AboutServer"
)

type actionAtom struct {
	DearFn func(writer http.ResponseWriter,s string,config MainTools.Config)	`json:"-"`
	Method string		`json:"method"`
	Description string	`json:"description"`
	Params []string		`json:"params"`
}
var actions []actionAtom = []actionAtom{
	{
		DearFn:      Apis.Api404Action,
		Method:      "api404",
		Description: "找不到指定接口",
		Params:		 []string{},
	},
	{
		DearFn:      Apis.GetRequestCodeAction,
		Method:      "getRequestCode",
		Description: "获取请求代码",
		Params:		 []string{},
	},
}
func InitAction() {
	http.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		writer.Header().Set("content-type", "application/json")             //返回数据格式是json
		params,err := AboutServer.PostParams(request)
		if err != nil {
			config.Logger.Printf("[action][fail] 解析失败")
			rObj := AboutServer.ReturnObj{}
			rObj.SetFail("解析参数发生时错误")
			rObj.Send(writer)
		} else {
			config.Logger.Printf("[action:api][Method:%s] 执行方法",params.Method)
			did := false
			for _,action := range actions {
				if action.Method == params.Method {
					did = true
					action.DearFn(writer,params.Content,config)
				}
			}
			if !did {
				Apis.Api404Action(writer,params.Content,config)
				config.Logger.Printf("[action][Method:%s] 找不到该方法",params.Method)
			}
		}
	})
	http.HandleFunc("/apis", func(writer http.ResponseWriter, request *http.Request) {
		config.Logger.Printf("[action:apis] 执行方法")
		rObj := AboutServer.ReturnObj{}
		rObj.SetContent(actions)
		rObj.Send(writer)
	})
}