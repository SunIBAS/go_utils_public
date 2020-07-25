package InitHttp

import (
	"encoding/json"
	"public.sunibas.cn/go_utils_public/src/utils/DirAndFile"
	"strings"
)
type Config struct {
	CwdPath string `json:"cwd_path"`
	DbPath string `json:"db_path"`
	Port string `json:"port"`
}
// 获取配置信息
// 并完成一系列初始化工作，包括以下内容
func Entrance(configPath string)(conf Config) {
	conf = Config{
		CwdPath: DirAndFile.GetCwd(),
		DbPath:  "",
	}
	if len(configPath) != 0 {
		json.Unmarshal([]byte(strings.Join(DirAndFile.ReadAsFileAsLine(configPath),"\n")),&conf)
	}
	conf.DbPath = conf.CwdPath + "\\db.db"
	conf.Port = ":3000"
	return
}
