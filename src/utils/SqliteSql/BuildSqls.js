if (process.argv.length < 3) {
    throw new Error("需要第二个参数表示输入json");
}
const dbJsonPath = process.argv[2];
//"C:\\Users\\HUZENGYUN\\go\\src\\sunibas.cn\\go_utils\\DirToDb\\DearLang\\db.json";
// const goFilePath = "C:\\Users\\HUZENGYUN\\go\\src\\sunibas.cn\\go_utils\\DirToDb\\DearLang\\Sqls.go";
let buildNode = {
    outputPath: "C:\\Users\\HUZENGYUN\\go\\src\\sunibas.cn\\go_utils\\DirToDb\\DearLang\\Sqls.go",
    packageName: "DearLang"
};

const fs = require('fs');
const db = require(dbJsonPath);

let tpl = `type {EntityName}Entity struct {
{{fields}}
}
func Get{EntityName}Table() SqliteSql.Table {
\t{entityName}Table := SqliteSql.Table{
\t\tName:   "{TableName}",
\t\tFields: []SqliteSql.Field{
{{tableField}}
\t\t},
\t\tCreateIndex:"{{CreateIndex}}",
\t\tNote: "{{note}}",
\t}
\treturn {entityName}Table
}
// 将查询结果转换为数组对象
func ParseRowBySelect{EntityName}(database * sql.DB,where string,table SqliteSql.Table) []{EntityName}Entity{
\treturn Query{EntityName}(table.GetSelectSql(where),database)
}
func Query{EntityName}(sql string,db * sql.DB) []{EntityName}Entity {
\tif rows,err := db.Query(sql);
\t\terr == nil {
\t\treturn ParseRows{EntityName}(rows)
\t} else {
\t\t// todo 错误暂时无解决方案
\t\tpanic(err)
\t}
}
func ParseRows{EntityName}(rows * sql.Rows) []{EntityName}Entity {
\t{entityName}s := []{EntityName}Entity{}
\tfor rows.Next() {
\t\t{entityName} := {EntityName}Entity{}
\t\trows.Scan(
{{scans}}
\t\t)
\t\t{entityName}s = append({entityName}s,{entityName})
\t}
\treturn {entityName}s
}

func (entity {EntityName}Entity)InsertOrUpdate(database * sql.DB,{entityName}Table SqliteSql.Table)  {
\tvalues,_ := SqliteSql.GetInsertValues(entity,{entityName}Table,entity.Id)
\tinsertOrUpdate := {entityName}Table.GetSpecialInsertSql("insertOrUpdate",values)
\tSqliteSql.ExecSqlString(database,insertOrUpdate)
}
`;
let out = ``;
db.forEach(table => {
    if ("no" in table) {
        if (table.tag == "build") {
            buildNode = table;
        }
        return;
    }
    let cp = tpl;
    cp = cp.replace(/{EntityName}/g,table.EntityName);
    cp = cp.replace(/{entityName}/g,table.entityName);
    cp = cp.replace(/{TableName}/g,table.TableName);
    let fields = "";
    let scans = "";
    if ("sql" in table) {
        if ("CreateIndex" in table.sql) {
            cp = cp.replace("{{CreateIndex}}",table.sql.CreateIndex);
        } else {
            cp = cp.replace("{{CreateIndex}}",'');
        }
    } else {
        cp = cp.replace("{{CreateIndex}}",'');
    }
    var tableField = [];
    table.fields.forEach(f => {
        let name = f.Name[0].toUpperCase() + f.Name.substring(1);
        fields += "\t" + name + " " + f._type + "\r\n";
        scans += "\t\t\t&" + table.entityName + "." + name + ",\r\n";
        let _f = {};
        for (let i in f) {
            if (i.substring(0,1) != "_") {
                _f[i] = f[i];
            }
        }
        _f.Name = _f.Name.toLowerCase();
        tableField.push(_f);
    });
    cp = cp.replace("{{fields}}",fields);
    cp = cp.replace("{{scans}}",scans);
    tableField = JSON.stringify(tableField,undefined,'\t');
    tableField = tableField.split('\n')
        .map(line => {
            if (line === ']' || line === '[') {
                return '';
            } else if (line === '\t{' || line === '\t},') {
                return "\t\t" + line;
            } else if (line === "\t}") {
                return "\t\t\t},";
            } else {
                line = line.split('');
                line[line.indexOf('"')] = '';
                line[line.indexOf('"')] = '';
                if (line[line.length - 1] !== ",") {
                    line.push(',')
                }
                return "\t\t" + line.join('');
            }
        })
        .filter(_ => _)
        .join('\r\n');
    cp = cp.replace("{{tableField}}",tableField);
    cp = cp.replace("{{note}}",table.note);
    out += cp + "\r\n\r\n";
});

let headOut = `package ${buildNode.packageName}

import (
\t"database/sql"
\t"fmt"
\t"public.sunibas.cn/go_utils_public/src/utils/DirAndFile"
\t"public.sunibas.cn/go_utils_public/src/utils/SqliteSql"
)

`;

let footTpl = `// 启动前调用该方法
// db,tables := Init("") // <- 保存好这两个变量，后面的数据库操作都需要用到他们
// defer db.Close()
func Init(dbPath string) (* sql.DB , map[string]SqliteSql.Table) {
\ttables := map[string]SqliteSql.Table{}
\tvar langTablesCreater map[string]func()SqliteSql.Table
\tif dbPath == "" {
\t\treturn nil,nil
\t} else {
\t\texist,err := DirAndFile.CheckTypeAndExist(dbPath)
\t\tif err == nil && exist != 1 {
\t\t} else if err != nil {
\t\t\tpanic(err)
\t\t} else {
\t\t\tfmt.Println("langDbPath 必须是文件")
\t\t}
\t\t_database,err := sql.Open("sqlite3", dbPath)
\t\tif err != nil {
\t\t\tpanic(err)
\t\t}

\t\tlangTablesCreater = map[string]func()SqliteSql.Table{
{{tables}}
\t\t}
\t\tfor k,v := range langTablesCreater {
\t\t\ttable := v()
\t\t\ttable.GetSqls()
\t\t\ttables[k] = table
\t\t\tesqls := []string {"create","createIndex"}
\t\t\tfor _,sqlName := range esqls {
\t\t\t\tif "" != tables[k].Sqls[sqlName] {
\t\t\t\t\tstatement,_ := _database.Prepare(tables[k].Sqls[sqlName])
\t\t\t\t\tstatement.Exec()
\t\t\t\t}
\t\t\t}
\t\t}

\t\treturn _database,tables
\t}
}`;

let tables = [];
db.forEach(table => {
    if ("no" in table) {
        return;
    }
    tables.push(`\t\t\t"${table.TableName}":Get${table.EntityName}Table,`);
});

out += footTpl.replace("{{tables}}",tables.join('\r\n'));
fs.writeFileSync(buildNode.outputPath,headOut + out,'utf-8');

