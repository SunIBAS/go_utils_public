package SqliteSql

import "strings"

// 字段信息，默认都是可空的，主键是默认不可空的
type Field struct {
	Name string
	Type string
	Note string
	Primary bool
	//notNull bool
}
// 表格信息
type Table struct {
	Name string
	Fields []Field
	multiKey []string
	Note string
	Sqls map[string] string
	CreateIndex string
}

// 获取一些 sql 便于后面的使用
// https://www.runoob.com/sqlite/sqlite-create-table.html
func (table * Table)GetSqls() {
	sqls := map[string] string {}
	createSql := "CREATE TABLE IF NOT EXISTS " + table.Name + " ("
	selectSql := "SELECT "
	insertSql := "INSERT INTO " + table.Name + " ("
	insertOrUpdateSql := "INSERT OR REPLACE INTO " + table.Name + " ("
	createFieldSql := []string{}
	insertFieldSql := []string{}
	insertFieldQue := []string{}
	selectTableNameDotFieldQue := []string{
		// "dir.id"
	}
	multiKey := ""
	Primary := false
	if table.multiKey != nil {
		Primary = true
		multiKey = ",Primary key (" + strings.Join(table.multiKey,",") + ")"
	}
	for _,f := range table.Fields {
		createField := "`" + f.Name + "` " + f.Type
		insertFieldSql = append(insertFieldSql,"`" + f.Name + "`")
		selectTableNameDotFieldQue = append(selectTableNameDotFieldQue,"`" + table.Name + "`.`" + f.Name + "`")
		insertFieldQue = append(insertFieldQue,"?")
		if f.Primary && !Primary {
			Primary = true
			createField += " PRIMARY KEY"
		}
		createFieldSql = append(createFieldSql,createField)

	}
	sqls["create"] = createSql + strings.Join(createFieldSql,",") + multiKey + ");"
	sqls["createIndex"] = table.CreateIndex
	sqls["insertStatement"] = insertSql + strings.Join(insertFieldSql,",") + ") VALUES (" + strings.Join(insertFieldQue,",") + ");"
	sqls["insert"] = insertSql + strings.Join(insertFieldSql,",") + ") VALUES ('$$');"
	sqls["insertOrUpdate"] = insertOrUpdateSql + strings.Join(insertFieldSql,",") + ") VALUES ('$$');"
	sqls["selectLeftJoin"] = selectSql + strings.Join(selectTableNameDotFieldQue,",") + " FROM " + table.Name
	sqls["select"] = selectSql + strings.Join(insertFieldSql,",") + "FROM " + table.Name
	table.Sqls = sqls
}

// 从上面 GetSqls 中生成一个 insert 语句
func (table * Table)GetInsertSql(values []string) string {
	return strings.Replace(table.Sqls["insert"],"$$",strings.Join(values,`','`),-1)
}
// 获取指定insert语句
func (table * Table)GetSpecialInsertSql(tar string,values []string) string {
	return strings.Replace(table.Sqls[tar],"$$",strings.Join(values,`','`),-1)
}

func (table * Table)GetSelectSql(where string) string {
	return table.Sqls["select"] + " " + where
}
