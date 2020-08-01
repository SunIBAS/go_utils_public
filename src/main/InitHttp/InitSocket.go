package InitHttp

import (
	"public.sunibas.cn/go_utils_public/src/main/Apis"
	"public.sunibas.cn/go_utils_public/src/utils/AboutServer"
)


func InitSocket() {
	AboutServer.InitSocket("/socket",ToDearSocketMessage)
}

func ToDearSocketMessage(msg []byte,l * AboutServer.Longsocket) error {
	sm,err := AboutServer.ToSM(msg)
	if err != nil {
		config.Logger.Println("[socket][fail] 解析失败")
		sm.SetFail("TO.SM FAIL")
		l.Write([]byte(err.Error()))
		return err
	} else {
		config.Logger.Printf("[socket][Method:%s] 正在执行方法",sm.Method)
		sm.Lsocket = l
		// todo 较为随意的在这里后面添加方法即可
		if sm.Method == "sayHello" {
			Apis.SayHelloSocket(sm,config)
		} else {
			config.Logger.Printf("[socket][Method:%s] 找不到指定方法",sm.Method)
			sm.SetFail("nothing").Return()
			sm.SetEnd("end").Return()
		}
		return nil
	}
}