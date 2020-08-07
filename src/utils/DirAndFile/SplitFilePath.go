package DirAndFile

import (
	"os"
	"path"
	"path/filepath"
)

func SplitPath(files string) (onlyPath string,onlyName string,ext string) {
	//files := "E:\\data\\test.txt"
	paths, _ := filepath.Split(files)
	onlyPath = paths      //获取路径中的目录及文件名 E:\data\  test.txt
	onlyName = filepath.Base(files) //获取路径中的文件名test.txt
	ext = path.Ext(files)      //获取路径中的文件的后缀 .txt
	return
}

func CreateFileBeforeCheckAndCreate(files string) {
	op,_,_ := SplitPath(files)
	os.MkdirAll(op,os.ModePerm)
}