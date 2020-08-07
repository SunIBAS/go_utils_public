package MainTools

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"public.sunibas.cn/go_utils_public/src/utils/SqliteSql"
)

type Config struct {
	// 日志
	Logger * log.Logger `json:"-"`
	// 数据库
	DB * sql.DB `json:"-"`
	// 表格
	Tables map[string]SqliteSql.Table `json:"-"`
	// 上面的内容不保存导出到 json 文件中

	CwdPath string `json:"cwd_path"`
	DbPath string `json:"db_path"`
	Port string `json:"port"`
	LogPath string `json:"log_path"`
	NodePath string `json:"node_path"`
}
