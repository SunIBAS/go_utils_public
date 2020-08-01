package AboutServer

import (
	"encoding/json"
)

type SocketMessage struct {
	Id string `json:"id"`
	Method string `json:"method"`
	Content string `json:"content"`
	Lsocket * Longsocket `json:"-"`
	Code int `json:"code"`
	Message string `json:"message"`
}

func (sm SocketMessage)ToBytes() (jsonStr []byte) {
	jsonStr,_ = json.MarshalIndent(sm,"","")
	return
}

func ToSM(msg []byte) (sm SocketMessage,err error){
	sm = SocketMessage{}
	err = json.Unmarshal(msg,&sm)
	//sm.Method = strings.ToLower(sm.Method)
	return
}

// 以下方法是为了和 RequestReturn 进行对应
func (sm * SocketMessage)SetContent(data interface{}) {
	jsonStr,_ := json.MarshalIndent(data,"","")
	sm.Content = string(jsonStr)
}
func (sm * SocketMessage)SetError(msg string) * SocketMessage {
	sm.Code = Error
	if len(msg) == 0 {
		sm.Message = "error"
	} else {
		sm.Message = msg
	}
	return sm
}
func (sm * SocketMessage)SetFail(msg string) * SocketMessage {
	sm.Code = Fail
	if len(msg) == 0 {
		sm.Message = "fail"
	} else {
		sm.Message = msg
	}
	return sm
}
func (sm * SocketMessage)SetSuccess(msg string) * SocketMessage {
	sm.Code = Success
	if len(msg) == 0 {
		sm.Message = "success"
	} else {
		sm.Message = msg
	}
	return sm
}
// 这里多余的一个方法是因为 socket 不会因为完成发送而中断，中断需要双方协作
func (sm * SocketMessage)SetEnd(msg string) * SocketMessage {
	sm.Code = End
	if len(msg) == 0 {
		sm.Message = "end"
	} else {
		sm.Message = msg
	}
	return sm
}
func (sm SocketMessage)Return() SocketMessage {
	sm.Lsocket.Write(sm.ToBytes())
	return sm
}