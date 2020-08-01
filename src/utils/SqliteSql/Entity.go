package SqliteSql

import (
	uuid "github.com/satori/go.uuid"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// 暂时没用到
type Entity interface {
	GetTable() Table
}

// 获取一个完整的插入语句，返回 sql ， id:uuid
func GetInsertSql(s interface{},tab Table) (sql string,uid string) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	kv := map[string] interface{}{}
	var err error
	uid = uuid.Must(uuid.NewV4(),err).String()
	for i := 0;i < t.NumField();i++ {
		kv[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
	}
	values := []string{}
	for _,f := range tab.Fields {
		if f.Name == "id" {
			values = append(values,uid)
		} else {
			var v = ""
			if f.Type == "TEXT"{
				//values = append(values,kv[f.Name].(string))
				v = kv[f.Name].(string)
				v = strings.Replace(v,"'","''",-1)
			} else if f.Type == "INTEGER" {
				//values = append(values,kv[f.Name].(int64))
				v = strconv.FormatInt(kv[f.Name].(int64),10)
			} else if f.Type == "REAL" {
				//values = append(values,kv[f.Name].(float64))
				v = strconv.FormatFloat(kv[f.Name].(float64), 'E', -1, 32)
			} else if f.Type == "BLOB" {
				v = kv[f.Name].(string)
				v = strings.Replace(v,"'","''",-1)
			}
			values = append(values,v)
		}
	}
	sql = tab.GetInsertSql(values)
	return
}
// 配合 table 自动生成的 sql ，将 实体 转化为 insert sql 的 values 部分，返回 values id
// 传入 id 一般用于更新
func GetInsertValues(s interface{},tab Table,id string) (values []string,u string) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	kv := map[string] interface{}{}
	var err error
	if id == "" {
		u = uuid.Must(uuid.NewV4(), err).String()
	} else {
		u = id
	}
	for i := 0;i < t.NumField();i++ {
		kv[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
	}
	values = []string{}
	for _,f := range tab.Fields {
		if f.Name == "id" {
			values = append(values,u)
		} else {
			var v = ""
			if f.Type == "TEXT"{
				//values = append(values,kv[f.Name].(string))
				v = kv[f.Name].(string)
			} else if f.Type == "INTEGER" {
				//values = append(values,kv[f.Name].(int64))
				v = strconv.FormatInt(kv[f.Name].(int64),10)
			} else if f.Type == "REAL" {
				//values = append(values,kv[f.Name].(float64))
				v = strconv.FormatFloat(kv[f.Name].(float64), 'E', -1, 32)
			}
			values = append(values,v)
		}
	}
	return
}

// https://blog.csdn.net/u011957758/article/details/81193806
// 表只有这两种类型 int64 string 的字段
// 将结构体转为字符串，用于输出日志方便
func EntityToString(s interface{}) string {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	out := ""
	for i := 0;i < t.NumField();i++ {
		ctype := t.Field(i).Type.Name()
		if ctype == "int64" {
			out += strings.ToLower(t.Field(i).Name) + " => " + strconv.FormatInt(v.Field(i).Interface().(int64),10) + "\r\n"
		} else if ctype == "string" {
			out += strings.ToLower(t.Field(i).Name) + " => " + v.Field(i).Interface().(string) + "\r\n"
		}
	}
	return out
}

// 获取一个时间戳
func GetTimestamp() int64 {
	return time.Now().Unix()
}
// 将时间戳转为 年月日-时分秒
func FormatTimestamp(ts int64) string {
	timeLayout := "2006-01-02 15:04:05"
	return time.Unix(ts,0).Format(timeLayout)
}
