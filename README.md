# go_utils_public

这个是从私有项目中单独分离出来的

> 本项目的 src 文件夹 📂 介绍请参见[此处](./docs/src.md)

> 使用到的技术及网址(以下网址尽可能保留非GitHub链接本地版本)

> 使用的 ui 为 [iview](https://www.iviewui.com/components)

[🔗 用Go实现CORS跨域资源共享的服务器支持](http://semicircle.github.io/blog/2013/09/29/go-with-cors/)

[🔗 statik 实现将html打包放入go中](https://github.com/rakyll/statik)
```shell script
# 本项目中的用法如下
statik -src=html
```

[🔗 websocket long connect](https://github.com/qianlnk/longsocket)

[🔗 编码识别](https://github.com/saintfish/chardet)

```go
detector：= chardet.NewTextDetector()
result，err：= detector.DetectBest(some_text)

if err == nil {
 fmt.Printf（
"检测到的字符集为％s，语言为％s ",
 result.Charset,
 result.Language）
} 
```

[🔗 Sqlite3](https://github.com/mattn/go-sqlite3)

```shell script
# 更新数据库代码的方法是
# 将工作目录移动到 BuildSql.js 目录下
# src/utils/SqliteSql>
node BuildSqls.js ./../../main/Sqls/db.json
```

----

iview 标签转换
```text
Button: i-button
Col: i-col
Table: i-table
Input: i-input
Form: i-form
Menu: i-menu
Select: i-select
Option: i-option
Progress: i-progress
Time: i-time
MenuItem：Menu-Item
RadioGroup：Radio-Group
<Icon / >：<Icon ></Icon>
FormItem：Form-Item
Header:i-header
Content:i-content
```