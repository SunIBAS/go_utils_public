package MySocket

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"net/url"
	"public.sunibas.cn/go_utils_public/src/utils/AboutServer"
)

var ToDearMessage func([]byte,* AboutServer.Longsocket) error
func InitSocket(path string,tdm func([]byte,* AboutServer.Longsocket) error) {
	ToDearMessage = tdm
	http.Handle(path, websocket.Handler(handleSocket))
}

func handleSocket(ws *websocket.Conn) {
	req := ws.Request()
	fmt.Println(req)
	//u, err := url.Parse(req.Header.Get("Origin"))
	_, err := url.Parse(req.Header.Get("Origin"))
	if err != nil {
		ws.Close()
		return
	}

	//user := u.Query().Get("user")
	//password := u.Query().Get("password")
	//fmt.Println(user, password)
	mysocket := AboutServer.NewConn("", "", "", false, 128*1024)
	mysocket.SetSocket(ws)
	defer mysocket.Close()
	go mysocket.WriteLoop()
	go mysocket.ReadLoop()
	mysocket.Read(func(bytes []byte, longsocket * AboutServer.Longsocket) error {
		return ToDearMessage(bytes,longsocket)
	})
}
