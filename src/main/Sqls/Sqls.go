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

func InsertOrUpdate(database * sql.DB,demoTable SqliteSql.Table,entity DemoEntity)  {
	values,_ := SqliteSql.GetInsertValues(entity,demoTable,entity.Id)
	insertOrUpdate := demoTable.GetSpecialInsertSql("insertOrUpdate",values)
	SqliteSql.ExecSqlString(database,insertOrUpdate)
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

/*
func InsertOrUpdate(database * sql.DB,pubwebTable SqliteSql.Table,entity PubWebEntity)  {
	values,_ := SqliteSql.GetInsertValues(entity,pubwebTable,entity.Id)
	insertOrUpdate := pubwebTable.GetSpecialInsertSql("insertOrUpdate",values)
	SqliteSql.ExecSqlString(database,insertOrUpdate)
}
*/


// 本质上讲，这个软件不应该包含 office 文档问题提取功能的，
// 或者说不应该是主体功能，但是鉴于这部分功能需要使用到上面定义的数据库
// 只能勉为其难的将这部分写到项目里面，其实不应该写进来的
// 另外一个声明就是，这部分的代码将直接使用 sql 不再写类似 modal 和 dao 的部分
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