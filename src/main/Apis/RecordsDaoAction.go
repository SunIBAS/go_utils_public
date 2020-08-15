package Apis

import (
	"encoding/json"
	"net/http"
	"public.sunibas.cn/go_utils_public/src/main/MainTools"
	"public.sunibas.cn/go_utils_public/src/main/Sqls"
	"public.sunibas.cn/go_utils_public/src/utils/AboutServer"
	"public.sunibas.cn/go_utils_public/src/utils/SqliteSql"
	"strconv"
	"strings"
	"time"
)

var recordTableName = "records"
var RecordsDaoActions []MainTools.ActionAtom = []MainTools.ActionAtom{
	{
		DearFn:      ListRecords,
		Method:      "ListRecords",
		Description: "获取记录块",
		Params:      []string {"tag","field1","field2","field3"},
	},
	{
		DearFn:      InsertRecords,
		Method:      "InsertRecords",
		Description: "插入记录块",
		Params:      []string {"tag","content","field1","field2","field3"},
	},
	{
		DearFn:      UpdateRecords,
		Method:      "UpdateRecords",
		Description: "更新",
		Params:      []string {"id","tag","content","field1","field2","field3"},
	},
	{
		DearFn:      GetByRecordsId,
		Method:      "GetByRecordsId",
		Description: "GetByRecordsId",
		Params:      []string {"id"},
	},
	{
		DearFn:      GetRecordsCount,
		Method:      "GetRecordsCount",
		Description: "获取records记录数",
		Params:      []string {},
	},
	{
		DearFn:      DeleteRecords,
		Method:      "DeleteRecords",
		Description: "删除一条记录",
		Params:      []string {"id"},
	},
}

// 不对content进行筛选
func ListRecords(writer http.ResponseWriter,s string,config MainTools.Config)  {
	where := ""
	rObj := AboutServer.ReturnObj{}
	var page MainTools.PageParam
	err := json.Unmarshal([]byte(s),&page)
	if len(page.Other) != 0 {
		var records Sqls.RecordsEntity
		err1 := json.Unmarshal([]byte(page.Other),&records)
		if err1 == nil {
			ws := []string{}
			if len(records.Tag) != 0 {
				ws = append(ws," `tag`='" + records.Tag + "' ")
			}
			if len(records.Field1) != 0 {
				ws = append(ws," `field1`='" + records.Field1 + "' ")
			}
			if len(records.Field2) != 0 {
				ws = append(ws," `field2`='" + records.Field2 + "' ")
			}
			if len(records.Field3) != 0 {
				ws = append(ws," `field3`='" + records.Field3 + "' ")
			}
			if len(ws) != 0 {
				where = " where " + strings.Join(ws,"and")
			}
		} else {
			err = err1
		}
	}
	if err != nil {
		rObj.SetFail("参数解析错误，err:" + err.Error()).Send(writer)
	} else {
		offset := page.Count * page.Page
		recordss := Sqls.ParseRowBySelectRecords(config.DB,where + " limit " + strconv.Itoa(page.Count) + " offset " + strconv.Itoa(offset),
			config.Tables[recordTableName])
		rObj.SetContent(recordss).SetSuccess("").Send(writer)
	}
}

func InsertRecords(writer http.ResponseWriter,s string,config MainTools.Config) {
	rObj := AboutServer.ReturnObj{}
	var records Sqls.RecordsEntity
	err := json.Unmarshal([]byte(s),&records)
	if err != nil {
		rObj.SetFail("参数解析错误，err:" + err.Error()).Send(writer)
	} else {
		records.CreateTime = strconv.FormatInt(time.Now().UnixNano(),10)
		sql,id := SqliteSql.GetInsertSql(records,config.Tables[recordTableName])
		SqliteSql.ExecSqlString(config.DB,sql)
		rObj.SetStringContent(id).SetSuccess("插入成功").Send(writer)
	}
}

// rec_s 是数据库的内容
// rec_d 是要修改的字段
func updateRecords(rec_s,rec_d Sqls.RecordsEntity) Sqls.RecordsEntity {
	if len(rec_d.Tag) != 0 { rec_s.Tag = rec_d.Tag }
	if len(rec_d.Content) != 0 { rec_s.Content = rec_d.Content }
	if len(rec_d.Bak1) != 0 { rec_s.Bak1 = rec_d.Bak1 }
	if len(rec_d.Bak2) != 0 { rec_s.Bak2 = rec_d.Bak2 }
	if len(rec_d.Field1) != 0 { rec_s.Field1 = rec_d.Field1 }
	if len(rec_d.Field2) != 0 { rec_s.Field2 = rec_d.Field2 }
	if len(rec_d.Field3) != 0 { rec_s.Field3 = rec_d.Field3 }
	return rec_d
}
func UpdateRecords(writer http.ResponseWriter,s string,config MainTools.Config) {
	rObj := AboutServer.ReturnObj{}
	var records Sqls.RecordsEntity
	err := json.Unmarshal([]byte(s),&records)
	if err != nil {
		rObj.SetFail("参数解析错误，err:" + err.Error()).Send(writer)
	} else {
		if records.Id != "" {
			recordss := Sqls.ParseRowBySelectRecords(config.DB," where id='" + records.Id + "'",config.Tables[recordTableName])
			if len(recordss) == 1 {
				records = updateRecords(recordss[0],records)
				records.InsertOrUpdate(config.DB,config.Tables[recordTableName])
				rObj.SetSuccess("插入成功").Send(writer)
			} else {
				rObj.SetFail("找不到对应记录").Send(writer)
			}
		} else {
			rObj.SetFail("id 为空").Send(writer)
		}
	}
}

func GetByRecordsId(writer http.ResponseWriter,s string,config MainTools.Config) {
	rObj := AboutServer.ReturnObj{}
	var records Sqls.RecordsEntity
	err := json.Unmarshal([]byte(s),&records)
	if err != nil {
		rObj.SetFail("参数解析错误，err:" + err.Error()).Send(writer)
	} else {
		recordss := Sqls.ParseRowBySelectRecords(config.DB," where id='" + records.Id + "'",config.Tables[recordTableName])
		rObj.SetContent(recordss).SetSuccess("").Send(writer)
	}
}

func GetRecordsCount(writer http.ResponseWriter,s string,config MainTools.Config) {
	count := SqliteSql.QueryCount(recordTableName,config.DB)
	rObj := AboutServer.ReturnObj{}
	rObj.SetStringContent(strconv.Itoa(count)).SetSuccess("").Send(writer)
}

func DeleteRecords(writer http.ResponseWriter,s string,config MainTools.Config) {
	rObj := AboutServer.ReturnObj{}
	var records Sqls.RecordsEntity
	err := json.Unmarshal([]byte(s),&records)
	if err != nil {
		rObj.SetFail("参数解析错误，err:" + err.Error()).Send(writer)
	} else {
		var sql = "delete from " + recordTableName + " where `id`='" + records.Id + "'"
		SqliteSql.ExecSqlString(config.DB,sql)
		rObj.SetStringContent("删除成功").SetSuccess("").Send(writer)
	}
}