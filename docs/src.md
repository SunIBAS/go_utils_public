## src 文件夹下的文件说明

- MySocket/SocketMessage.go 使用方法

```go
package _demo_
import (
    "public.sunibas.cn/go_utils_public/src/MySocket"
	"public.sunibas.cn/go_utils_public/src/utils/AboutServer"
)
func dear(msg []byte,l * MyServer.Longsocket) {
    sm,err := MySocket.ToSM(msg)
    if err != nil {
        // 发送错误信息
        sm.SetFail(err.Error())
        sm.Send()
        // 结束 socket 通信
        sm.SetEnd()
        sm.Send()
    } else {
        // 正常通信
    }
}
```

- utils/AboutServer/GetParams.go

```go
package _demo_
import (
	"net/http"
	"public.sunibas.cn/go_utils_public/src/utils/AboutServer"
)
func test1() {
    http.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
        params,err := AboutServer.PostParams(request)
        if err != nil {
        } else {
        	// params.Method ...
        }
}
```

- utils/AboutServer/longsocket.go

```见 GitHub ```

- utils/AboutServer/RequestReturn.go

```go
rObj := Utils.ReturnObj{}
rObj....
```

- 📂utils/DirAndFile

```text
cwdPath.go 获取执行目录
```

