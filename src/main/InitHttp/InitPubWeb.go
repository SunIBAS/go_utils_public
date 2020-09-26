package InitHttp

import (
	"encoding/json"
	"net/http"
	"public.sunibas.cn/go_utils_public/src/main/Sqls"
	"public.sunibas.cn/go_utils_public/src/utils/AboutServer"
	"public.sunibas.cn/go_utils_public/src/utils/DirAndFile"
	"public.sunibas.cn/go_utils_public/src/utils/SqliteSql"
	"strconv"
	"time"
)

// 这里的目的是发布一个本地的网站，访问的根部是 pubWeb

var (
	pubWebRoot = "/pubWeb"
	// {url:run} {链接:是否启动}
	pubWebMap = map[string] bool{}
	pubWebTableName = "pubweb"
)

func InitPubWeb() {
	http.HandleFunc(pubWebRoot + "Action", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		writer.Header().Set("content-type", "application/json")             //返回数据格式是json

		params,err := AboutServer.PostParams(request)
		rObj := AboutServer.ReturnObj{ }
		if err == nil {
			if "add" == params.Method {
				pubWebAdd(params,rObj,writer)
			} else if "list" == params.Method {
				pubWebList(params,rObj,writer)
			} else if "checkUrl" == params.Method {
				pubWebCheckUrl(params,rObj,writer)
			} else if "run" == params.Method {
				pubWebRun(params,rObj,writer)
			} else if "delete" == params.Method {
				pubWebDelete(params,rObj,writer)
			}
		}
	})

	pubwebs := Sqls.ParseRowBySelectPubWeb(config.DB,"",config.Tables[pubWebTableName])
	for _,pw := range pubwebs {
		pubWebMap[pubWebRoot + "/" + pw.PerPath] = false
	}
}

func pubWebAdd(params AboutServer.Parmas,rObj AboutServer.ReturnObj,writer http.ResponseWriter) {
	var pubweb Sqls.PubWebEntity
	perr := json.Unmarshal([]byte(params.Content),&pubweb)
	if perr == nil {
		if t,err := DirAndFile.CheckTypeAndExist(pubweb.Dir);err == nil && t == 1 {
			if _,ok := pubWebMap[pubWebRoot + "/" + pubweb.PerPath]; ok {
				// 已经注册过了
				rObj.SetFail("链接已被使用").Send(writer)
			} else {
				sql,_ := SqliteSql.GetInsertSql(Sqls.PubWebEntity{
					Id:         "",
					Dir:        pubweb.Dir,
					PerPath:    pubweb.PerPath,
					Type:       "web",
					Status:     "start",
					CreateTime: strconv.FormatInt(time.Now().UnixNano(),10),
				},config.Tables[pubWebTableName])
				SqliteSql.ExecSqlString(config.DB,sql,config.Logger)
				pubWebMap[pubWebRoot + "/" + pubweb.PerPath] = false
				rObj.SetSuccess("定义成功,链接为 :" + pubWebRoot + "/" + pubweb.PerPath).Send(writer)
			}
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
	} else {
		rObj.SetFail("解析请求参数").Send(writer)
	}
}

func pubWebList(params AboutServer.Parmas,rObj AboutServer.ReturnObj,writer http.ResponseWriter)  {
	//var page MainTools.PageParam
	//err := json.Unmarshal([]byte(params.Content),&page)
	//if err != nil {
	//	rObj.SetFail("参数解析错误，err:" + err.Error()).Send(writer)
	//} else {
		//offset := page.Count * page.Page
		//pubs := Sqls.ParseRowBySelectPubWeb(config.DB," limit " + strconv.Itoa(page.Count) + " offset " + strconv.Itoa(offset),
		//	config.Tables[pubWebTableName])
		pubs := Sqls.ParseRowBySelectPubWeb(config.DB,"", config.Tables[pubWebTableName])
		rObj.SetContent(pubs).SetSuccess("").Send(writer)
	//}
}

func pubWebCheckUrl(params AboutServer.Parmas,rObj AboutServer.ReturnObj,writer http.ResponseWriter) {
	var pubweb Sqls.PubWebEntity
	perr := json.Unmarshal([]byte(params.Content),&pubweb)
	if perr != nil {
		rObj.SetFail("参数解析错误").Send(writer)
	} else {
		if _,ok := pubWebMap[pubWebRoot + "/" + pubweb.PerPath]; ok {
			// 已经注册过了
			rObj.SetStringContent("fail").SetSuccess("链接已被使用").Send(writer)
		} else {
			rObj.SetStringContent("ok").SetSuccess("链接可以使用").Send(writer)
		}
	}
}

func pubWebRun(params AboutServer.Parmas,rObj AboutServer.ReturnObj,writer http.ResponseWriter)  {
	var pubweb Sqls.PubWebEntity
	perr := json.Unmarshal([]byte(params.Content),&pubweb)
	if perr != nil {
		rObj.SetFail("参数解析错误").Send(writer)
	} else {
		if _,ok := pubWebMap[pubWebRoot + "/" + pubweb.PerPath]; ok {
			if pubWebMap[pubWebRoot + "/" + pubweb.PerPath] == false {
				http.Handle(pubWebRoot + "/" + pubweb.PerPath + "/",
					http.StripPrefix(pubWebRoot + "/" + pubweb.PerPath + "/",
						http.FileServer(http.Dir(pubweb.Dir))))
				pubWebMap[pubWebRoot + "/" + pubweb.PerPath] = true
			}
		}
		rObj.SetStringContent("ok").SetSuccess("启动成功").Send(writer)
	}
}

func pubWebDelete(params AboutServer.Parmas,rObj AboutServer.ReturnObj,writer http.ResponseWriter) {
	var pubweb Sqls.PubWebEntity
	perr := json.Unmarshal([]byte(params.Content),&pubweb)
	if perr != nil {
		rObj.SetError(perr.Error()).Send(writer)
	} else {
		var sql = "delete from " + pubWebTableName + " where `id`='" + pubweb.Id + "'"
		SqliteSql.ExecSqlString(config.DB,sql,config.Logger)
		delete(pubWebMap,pubWebRoot + "/" + pubweb.PerPath)
		rObj.SetSuccess("删除成功").Send(writer)
	}
}