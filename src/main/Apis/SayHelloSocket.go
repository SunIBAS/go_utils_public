package Apis

import (
	"public.sunibas.cn/go_utils_public/src/main/MainTools"
	"public.sunibas.cn/go_utils_public/src/utils/AboutServer"
)

func SayHelloSocket(sm AboutServer.SocketMessage,config MainTools.Config) {
	// 例如解析一个对象
	//var demo Sqls.DemoEntity
	//err := json.Unmarshal([]byte(sm.Content),&demo)
	sm.SetSuccess("hello").Return()

	sm.SetEnd("end").Return()
}

