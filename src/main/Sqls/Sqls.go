package Sqls

import (
	"database/sql"
	"fmt"
	"public.sunibas.cn/go_utils_public/src/utils/DirAndFile"
	"public.sunibas.cn/go_utils_public/src/utils/SqliteSql"
)

type DemoEntity struct {
	Id string
	Lang string
	Description string
	Content string
	CreateTime string
	Tags string
	Bak1 string
	Bak2 string

}
func GetDemoTable() SqliteSql.Table {
	demoTable := SqliteSql.Table{
		Name:   "demo",
		Fields: []SqliteSql.Field{
			{
				Name: "id",
				Type: "TEXT",
				Primary: true,
				Note: "",
			},
			{
				Name: "lang",
				Type: "TEXT",
				Note: "代码语言",
			},
			{
				Name: "description",
				Type: "TEXT",
				Note: "描述",
			},
			{
				Name: "content",
				Type: "TEXT",
				Note: "内容",
			},
			{
				Name: "createtime",
				Type: "TEXT",
				Note: "创建时间",
			},
			{
				Name: "tags",
				Type: "TEXT",
				Note: "标签",
			},
			{
				Name: "bak1",
				Type: "TEXT",
				Note: "备用1",
			},
			{
				Name: "bak2",
				Type: "TEXT",
				Note: "备用2",
			},
		},
		CreateIndex:"CREATE INDEX if not exists 'demo_id' ON 'demo' ( 'id' ASC );",
		Note: "代码",
	}
	return demoTable
}
// 将查询结果转换为数组对象
func ParseRowBySelectDemo(database * sql.DB,where string,table SqliteSql.Table) []DemoEntity{
	return QueryDemo(table.GetSelectSql(where),database)
}
func QueryDemo(sql string,db * sql.DB) []DemoEntity {
	if rows,err := db.Query(sql);
		err == nil {
		return ParseRowsDemo(rows)
	} else {
		// todo 错误暂时无解决方案
		panic(err)
	}
}
func ParseRowsDemo(rows * sql.Rows) []DemoEntity {
	demos := []DemoEntity{}
	for rows.Next() {
		demo := DemoEntity{}
		rows.Scan(
			&demo.Id,
			&demo.Lang,
			&demo.Description,
			&demo.Content,
			&demo.CreateTime,
			&demo.Tags,
			&demo.Bak1,
			&demo.Bak2,

		)
		demos = append(demos,demo)
	}
	return demos
}

func (entity DemoEntity)InsertOrUpdate(database * sql.DB,demoTable SqliteSql.Table)  {
	values,_ := SqliteSql.GetInsertValues(entity,demoTable,entity.Id)
	insertOrUpdate := demoTable.GetSpecialInsertSql("insertOrUpdate",values)
	SqliteSql.ExecSqlString(database,insertOrUpdate,nil)
}


type PubWebEntity struct {
	Id string
	Dir string
	PerPath string
	Type string
	Status string
	CreateTime string
	Bak1 string
	Bak2 string

}
func GetPubWebTable() SqliteSql.Table {
	pubwebTable := SqliteSql.Table{
		Name:   "pubweb",
		Fields: []SqliteSql.Field{
			{
				Name: "id",
				Type: "TEXT",
				Primary: true,
				Note: "",
			},
			{
				Name: "dir",
				Type: "TEXT",
				Note: "文件路径",
			},
			{
				Name: "perpath",
				Type: "TEXT",
				Note: "网站前缀",
			},
			{
				Name: "type",
				Type: "TEXT",
				Note: "服务类型，例如网站，文件服务",
			},
			{
				Name: "status",
				Type: "TEXT",
				Note: "状态（启用、弃用）",
			},
			{
				Name: "createtime",
				Type: "TEXT",
				Note: "创建时间",
			},
			{
				Name: "bak1",
				Type: "TEXT",
				Note: "备用1",
			},
			{
				Name: "bak2",
				Type: "TEXT",
				Note: "备用2",
			},
		},
		CreateIndex:"",
		Note: "发布本地文件作为一个网站",
	}
	return pubwebTable
}
// 将查询结果转换为数组对象
func ParseRowBySelectPubWeb(database * sql.DB,where string,table SqliteSql.Table) []PubWebEntity{
	return QueryPubWeb(table.GetSelectSql(where),database)
}
func QueryPubWeb(sql string,db * sql.DB) []PubWebEntity {
	if rows,err := db.Query(sql);
		err == nil {
		return ParseRowsPubWeb(rows)
	} else {
		// todo 错误暂时无解决方案
		panic(err)
	}
}
func ParseRowsPubWeb(rows * sql.Rows) []PubWebEntity {
	pubwebs := []PubWebEntity{}
	for rows.Next() {
		pubweb := PubWebEntity{}
		rows.Scan(
			&pubweb.Id,
			&pubweb.Dir,
			&pubweb.PerPath,
			&pubweb.Type,
			&pubweb.Status,
			&pubweb.CreateTime,
			&pubweb.Bak1,
			&pubweb.Bak2,

		)
		pubwebs = append(pubwebs,pubweb)
	}
	return pubwebs
}

func (entity PubWebEntity)InsertOrUpdate(database * sql.DB,pubwebTable SqliteSql.Table)  {
	values,_ := SqliteSql.GetInsertValues(entity,pubwebTable,entity.Id)
	insertOrUpdate := pubwebTable.GetSpecialInsertSql("insertOrUpdate",values)
	SqliteSql.ExecSqlString(database,insertOrUpdate,nil)
}


type RecordsEntity struct {
	Id string
	Tag string
	Content string
	CreateTime string
	Field1 string
	Field2 string
	Field3 string
	Bak1 string
	Bak2 string

}
func GetRecordsTable() SqliteSql.Table {
	recordsTable := SqliteSql.Table{
		Name:   "records",
		Fields: []SqliteSql.Field{
			{
				Name: "id",
				Type: "TEXT",
				Primary: true,
				Note: "",
			},
			{
				Name: "tag",
				Type: "TEXT",
				Note: "标记（可以用来区分存储的内容是什么）",
			},
			{
				Name: "content",
				Type: "TEXT",
				Note: "存储内容(base64)",
			},
			{
				Name: "createtime",
				Type: "TEXT",
				Note: "创建时间",
			},
			{
				Name: "field1",
				Type: "TEXT",
				Note: "字段一",
			},
			{
				Name: "field2",
				Type: "TEXT",
				Note: "字段二",
			},
			{
				Name: "field3",
				Type: "TEXT",
				Note: "字段三",
			},
			{
				Name: "bak1",
				Type: "TEXT",
				Note: "备用1",
			},
			{
				Name: "bak2",
				Type: "TEXT",
				Note: "备用2",
			},
		},
		CreateIndex:"CREATE INDEX if not exists 'records_id' ON 'records' ( 'id' ASC );",
		Note: "存储任意内容",
	}
	return recordsTable
}
// 将查询结果转换为数组对象
func ParseRowBySelectRecords(database * sql.DB,where string,table SqliteSql.Table) []RecordsEntity{
	return QueryRecords(table.GetSelectSql(where),database)
}
func QueryRecords(sql string,db * sql.DB) []RecordsEntity {
	if rows,err := db.Query(sql);
		err == nil {
		return ParseRowsRecords(rows)
	} else {
		// todo 错误暂时无解决方案
		panic(err)
	}
}
func ParseRowsRecords(rows * sql.Rows) []RecordsEntity {
	recordss := []RecordsEntity{}
	for rows.Next() {
		records := RecordsEntity{}
		rows.Scan(
			&records.Id,
			&records.Tag,
			&records.Content,
			&records.CreateTime,
			&records.Field1,
			&records.Field2,
			&records.Field3,
			&records.Bak1,
			&records.Bak2,

		)
		recordss = append(recordss,records)
	}
	return recordss
}

func (entity RecordsEntity)InsertOrUpdate(database * sql.DB,recordsTable SqliteSql.Table)  {
	values,_ := SqliteSql.GetInsertValues(entity,recordsTable,entity.Id)
	insertOrUpdate := recordsTable.GetSpecialInsertSql("insertOrUpdate",values)
	fmt.Println("[InsertOrUpdate]>>> " + insertOrUpdate)
	SqliteSql.ExecSqlString(database,insertOrUpdate,nil)
}

// 启动前调用该方法
// db,tables := Init("") // <- 保存好这两个变量，后面的数据库操作都需要用到他们
// defer db.Close()
func Init(dbPath string) (* sql.DB , map[string]SqliteSql.Table) {
	tables := map[string]SqliteSql.Table{}
	var langTablesCreater map[string]func()SqliteSql.Table
	if dbPath == "" {
		return nil,nil
	} else {
		exist,err := DirAndFile.CheckTypeAndExist(dbPath)
		if err == nil && exist != 1 {
		} else if err != nil {
			panic(err)
		} else {
			fmt.Println("langDbPath 必须是文件")
		}
		_database,err := sql.Open("sqlite3", dbPath)
		if err != nil {
			panic(err)
		}

		langTablesCreater = map[string]func()SqliteSql.Table{
			"demo":GetDemoTable,
			"pubweb":GetPubWebTable,
			"records":GetRecordsTable,
		}
		for k,v := range langTablesCreater {
			table := v()
			table.GetSqls()
			tables[k] = table
			esqls := []string {"create","createIndex"}
			for _,sqlName := range esqls {
				if "" != tables[k].Sqls[sqlName] {
					statement,_ := _database.Prepare(tables[k].Sqls[sqlName])
					statement.Exec()
				}
			}
		}

		return _database,tables
	}
}