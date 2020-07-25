package main

import (
	"fmt"
	"log"
	"os"
)

// 初始化日志
func InitLog(logPath string) * log.Logger {
	fmt.Println("logPath is ",logPath)
	file, err := os.Create(logPath)
	if err != nil {
		log.Fatalln("fail to create test.log file!")
	}
	logger := log.New(file, "", log.Llongfile)
	logger.SetFlags(log.LstdFlags)
	return logger
}
