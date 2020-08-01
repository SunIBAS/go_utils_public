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

- 📂 utils/DirAndFile

```text
cwdPath.go 获取执行目录
```

- 📂 main/Apis

```text
这里记录的是请求处理方法，普通 http 请求方法，文件名以 action 结尾
例如
Api404Action.go
格式如下（可以不限于）
其中的 config 提供了包括【日志】【数据库】等操作
func Api404Action(writer http.ResponseWriter,s string,config InitHttp.Config) {
	rObj := AboutServer.ReturnObj{}
	rObj.SetFail("找不到指定接口").Send(writer)
}

socket 请求方法以 socket 结尾
例如
SayHelloSocket.go
其中 config 同上，而 sm 是一个 socket 的主要内容，可以反复进行消息的传递
func SayHelloSocket(sm AboutServer.SocketMessage,config InitHttp.Config) {
	// 例如解析一个对象
	//var demo Sqls.DemoEntity
	//err := json.Unmarshal([]byte(sm.Content),&demo)
	sm.SetSuccess("hello").Return()

	sm.SetEnd("end").Return()
}
```

- 📂 main/InitHttp

| 文件 | 说明 |
|----|----|
| InitHttp | 复制初始化整个应用，将调用下面三个文件 |
| Entrance | 复制初始化数据库、日志和其他基础配置 |
| InitAction | 初始化请求，并引入 Apis 文件夹中 xxxAction |
| InitSocket | 初始化socket，并引入 Apis 文件夹中 xxxSocket |

```text
InitHttp 提供该目录中对外的入口方法 InitHttp，需要一个配置文件路径，或者一个 空路径
InitHttp("C:\\Users\\config.json")
或
InitHttp("")


Entrance 调用后将返回一个 Config 对象，用于调用 日志和数据库等工具
conf := Entrance("C:\\Users\\config.json")
或
conf := Entrance("")


InitAction 提供一个初始化请求的方法
InitAction()  //不需要任何参数
需要将定义在 src/main/Apis 下的 xxxAction 的定义完善到
actions 数组中
{
    DearFn:      Apis.Api404Action, // 方法
    Method:      "api404",          // 名称
    Description: "找不到指定接口",   // 功能
    Params:		 []string{},        // 参数列表
}


InitSocket 考虑到 socket 通信的内容可能较少，
这里再 src/main/Apis 下的 xxxSocket 直接以 if else if 的形式进行调用即可
// todo 较为随意的在这里后面添加方法即可
if sm.Method == "sayHello" {
    Apis.SayHelloSocket(sm,config)
}
```

- 📂 main/Sqls

```text
db.json 记录了要初始化的数据库信息，其中最后一个部分记录以下内容
{
    "no": true,
    // 这个是标志
    "tag": "build",
    // 这个是要生成 Sqls 文件的绝对路径
    "outputPath": "C:\\Users\\HUZENGYUN\\go\\src\\public.sunibas.cn\\go_utils_public\\src\\main\\Sqls\\Sqls.go",
    // 这个是包名，目前为 Sqls
    "packageName": "Sqls"
}
表的定义可以参考 demo 表的完整定义


Sqls.go 这个文件是自动生成的
生成方法是在 src/utils/SqliteSql 目录下执行
node BuildSqls.js ./../../main/Sqls/db.json
```



