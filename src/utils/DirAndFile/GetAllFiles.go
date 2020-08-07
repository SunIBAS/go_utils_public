package DirAndFile

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type FileNode struct {
	Name string
	File bool
	FullPath string
	Child []FileNode
}
// 路径分隔符，windows下应该是 \
const Filepath_Separator = string(rune(filepath.Separator))
// 仅仅获取当前一个文件夹下的 文件和文件夹
func GetSubDirOrFile(dirPath string) []FileNode {
	fileNode := []FileNode{}
	flist, e := ioutil.ReadDir(dirPath)
	if e != nil {
		fmt.Println("read ",dirPath)
		fmt.Println("read file error")
		return fileNode
	}
	for _, f := range flist {
		fileNode = append(fileNode, FileNode{
			Name: f.Name(),
			File: !f.IsDir(),
			FullPath: dirPath + Filepath_Separator + f.Name(),
		})
	}
	return fileNode
}
func GetAllSubDirOrFile(dirPath string) []FileNode {
	fileNodes := GetSubDirOrFile(dirPath)
	for ind,fn := range fileNodes {
		if !fn.File {
			fileNodes[ind].Child = GetAllSubDirOrFile(fn.FullPath)
		}
	}
	return fileNodes
}

func GetAllSubDirOrFileNotZip(dirPath string) []FileNode {
	fileNodes := GetSubDirOrFile(dirPath)
	out := []FileNode{}
	out = append(out,fileNodes...)
	for _,fn := range fileNodes {
		if !fn.File {
			out = append(out,GetAllSubDirOrFileNotZip(fn.FullPath)...)
		}
	}
	return out
}