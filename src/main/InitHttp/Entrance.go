package InitHttp

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"public.sunibas.cn/go_utils_public/src/main/MainTools"
	"public.sunibas.cn/go_utils_public/src/main/Sqls"
	"public.sunibas.cn/go_utils_public/src/utils/DirAndFile"
	"strings"
)

// 获取配置信息
// 并完成一系列初始化工作，包括以下内容
func Entrance(configPath string)(conf MainTools.Config) {
	conf = MainTools.Config{}
	if len(configPath) != 0 {
		json.Unmarshal([]byte(strings.Join(DirAndFile.ReadAsFileAsLine(configPath),"\n")),&conf)
	}
	conf.CwdPath = DirAndFile.GetCwd()
	if len(conf.DbPath) == 0 {
		conf.DbPath = conf.CwdPath + "\\db.db"
	}
	if len(conf.LogPath) == 0 {
		conf.LogPath = conf.CwdPath + "\\log.log"
	}
	if len(conf.NodePath) == 0 {
		conf.NodePath = conf.CwdPath + "\\node.exe"
	}
	conf.Logger = initLog(conf.LogPath)
	conf.Port = ":3000"
	conf.DB,conf.Tables = Sqls.Init(conf.DbPath)
	return
}

// 初始化日志
func initLog(logPath string) * log.Logger {
	fmt.Println("logPath is ",logPath)
	file, err := os.Create(logPath)
	if err != nil {
		log.Fatalln("fail to create test.log file!")
	}
	logger := log.New(file, "", log.Llongfile)
	logger.SetFlags(log.LstdFlags)
	return logger
}