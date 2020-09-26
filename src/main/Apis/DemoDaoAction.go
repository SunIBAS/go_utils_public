package Apis

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"public.sunibas.cn/go_utils_public/src/Datas"
	"public.sunibas.cn/go_utils_public/src/main/MainTools"
	"public.sunibas.cn/go_utils_public/src/main/Sqls"
	"public.sunibas.cn/go_utils_public/src/utils/AboutServer"
	"public.sunibas.cn/go_utils_public/src/utils/Console"
	"public.sunibas.cn/go_utils_public/src/utils/DirAndFile"
	"public.sunibas.cn/go_utils_public/src/utils/SqliteSql"
	"strconv"
	"strings"
	"time"
)

var demoTableName = "demo"
var DemoDaoActions []MainTools.ActionAtom = []MainTools.ActionAtom{
	{
		DearFn:      InsertDemo,
		Method:      "InsertDemo",
		Description: "插入代码块",
		Params:      []string {"lang","description","content"},
	},
	{
		DearFn:      ListDemo,
		Method:      "ListDemo",
		Description: "获取代码",
		Params:      []string {"page","count","other"},
	},
	{
		DearFn:      UpdateDemo,
		Method:      "UpdateDemo",
		Description: "更新",
		Params:      []string {"id","lang","description","content"},
	},
	{
		DearFn:      GetByDemoId,
		Method:      "GetByDemoId",
		Description: "GetByDemoId",
		Params:      []string {"id"},
	},
	{
		DearFn:      GetDemoCount,
		Method:      "GetDemoCount",
		Description: "获取demo记录数",
		Params:      []string {},
	},
	{
		DearFn:      DeleteDemo,
		Method:      "DeleteDemo",
		Description: "删除一条记录",
		Params:      []string {"id"},
	},
	{
		DearFn:      RunDeom,
		Method:      "RunDeom",
		Description: "运行代码",
		Params:      []string {"id"},
	},
}

func ListDemo(writer http.ResponseWriter,s string,config MainTools.Config)  {
	where := ""
	rObj := AboutServer.ReturnObj{}
	var page MainTools.PageParam
	err := json.Unmarshal([]byte(s),&page)
	if len(page.Other) != 0 {
		var demo Sqls.DemoEntity
		err1 := json.Unmarshal([]byte(page.Other),&demo)
		if err1 == nil {
			if len(demo.Lang) != 0 {
				where = "where `lang`='" + demo.Lang + "'"
			}
			if len(demo.Description) != 0 {
				if len(where) != 0 {
					where += " and `description` like '%" + demo.Description + "%'"
				} else {
					where = "where `description` like '%" + demo.Description + "%'"
				}
			}
		} else {
			err = err1
		}
	}
	if err != nil {
		rObj.SetFail("参数解析错误，err:" + err.Error()).Send(writer)
	} else {
		offset := page.Count * page.Page
		demos := Sqls.ParseRowBySelectDemo(config.DB,where + " limit " + strconv.Itoa(page.Count) + " offset " + strconv.Itoa(offset),
			config.Tables[demoTableName])
		rObj.SetContent(demos).SetSuccess("").Send(writer)
	}
}

func InsertDemo(writer http.ResponseWriter,s string,config MainTools.Config) {
	rObj := AboutServer.ReturnObj{}
	var demo Sqls.DemoEntity
	err := json.Unmarshal([]byte(s),&demo)
	if err != nil {
		rObj.SetFail("参数解析错误，err:" + err.Error()).Send(writer)
	} else {
		demo.CreateTime = strconv.FormatInt(time.Now().UnixNano(),10)
		demo.Lang = strings.ToUpper(demo.Lang)
		sql,_ := SqliteSql.GetInsertSql(demo,config.Tables[demoTableName])
		SqliteSql.ExecSqlString(config.DB,sql,config.Logger)
		rObj.SetSuccess("插入成功").Send(writer)
	}
}

func UpdateDemo(writer http.ResponseWriter,s string,config MainTools.Config) {
	rObj := AboutServer.ReturnObj{}
	var demo Sqls.DemoEntity
	err := json.Unmarshal([]byte(s),&demo)
	if err != nil {
		rObj.SetFail("参数解析错误，err:" + err.Error()).Send(writer)
	} else {
		if demo.Id != "" {
			demos := Sqls.ParseRowBySelectDemo(config.DB," where id='" + demo.Id + "'",config.Tables[demoTableName])
			if len(demos) == 1 {
				demo.InsertOrUpdate(config.DB,config.Tables[demoTableName])
				rObj.SetSuccess("插入成功").Send(writer)
			} else {
				rObj.SetFail("找不到对应记录").Send(writer)
			}
		} else {
			rObj.SetFail("id 为空").Send(writer)
		}
	}
}

func GetByDemoId(writer http.ResponseWriter,s string,config MainTools.Config) {
	rObj := AboutServer.ReturnObj{}
	var demo Sqls.DemoEntity
	err := json.Unmarshal([]byte(s),&demo)
	if err != nil {
		rObj.SetFail("参数解析错误，err:" + err.Error()).Send(writer)
	} else {
		demos := Sqls.ParseRowBySelectDemo(config.DB," where id='" + demo.Id + "'",config.Tables[demoTableName])
		rObj.SetContent(demos).SetSuccess("").Send(writer)
	}
}

func GetDemoCount(writer http.ResponseWriter,s string,config MainTools.Config) {
	count := SqliteSql.QueryCount(demoTableName,config.DB)
	rObj := AboutServer.ReturnObj{}
	rObj.SetStringContent(strconv.Itoa(count)).SetSuccess("").Send(writer)
}

func DeleteDemo(writer http.ResponseWriter,s string,config MainTools.Config) {
	rObj := AboutServer.ReturnObj{}
	var demo Sqls.DemoEntity
	err := json.Unmarshal([]byte(s),&demo)
	if err != nil {
		rObj.SetFail("参数解析错误，err:" + err.Error()).Send(writer)
	} else {
		var sql = "delete from " + demoTableName + " where `id`='" + demo.Id + "'"
		SqliteSql.ExecSqlString(config.DB,sql,config.Logger)
		rObj.SetStringContent("删除成功").SetSuccess("").Send(writer)
	}
}

func RunDeom(writer http.ResponseWriter,s string,config MainTools.Config) {
	rObj := AboutServer.ReturnObj{}
	var demo Sqls.DemoEntity
	err := json.Unmarshal([]byte(s),&demo)
	if err != nil {
		rObj.SetFail("参数解析错误，err:" + err.Error()).Send(writer)
	} else {
		demos := Sqls.ParseRowBySelectDemo(config.DB," where id='" + demo.Id + "'",config.Tables[demoTableName])
		if len(demos) == 1 {
			demo = demos[0]
			demo.Lang = strings.ToLower(demo.Lang)
			if demo.Lang == "nodejs" {
				ret,err := runNJ(demo,config)
				if err != nil {
					rObj.SetSuccess(err.Error()).Send(writer)
				} else {
					rObj.SetSuccess(ret).Send(writer)
				}
			} else {
				rObj.SetFail("此类型代码不能执行").Send(writer)
			}
		} else {
			rObj.SetFail("找不到指定代码").Send(writer)
		}
	}
}
func runNJ(demo Sqls.DemoEntity,config MainTools.Config) (ret string,err error) {
	filename := strconv.FormatInt(time.Now().UnixNano(),10) + ".js"
	fullPath := config.CwdPath + "tmp" + DirAndFile.Filepath_Separator + filename
	content := Datas.DecodeBase64AndURI(demo.Content)
	if len(content) != 0 {
		DirAndFile.WriteWithIOUtil(fullPath,content)
		ret,err = Console.RunNodejs(fullPath,config.NodePath)
		lines := strings.Split(ret,"\n")
		lines = lines[3:]
		lastLine := lines[len(lines) - 1]
		ret = ""
		for _,l := range lines {
			ret += strings.Replace(l,lastLine,">>>",1) + "\n"
		}
		ret = Datas.EncodeBase64(ret)
	} else {
		err = errors.New("获取到的脚本内容为空")
	}
	// todo 测试(如果需要调试可以加注释下一句)
	os.Remove(fullPath)
	return
}