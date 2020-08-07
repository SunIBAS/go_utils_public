package Apis

import (
	"encoding/json"
	"net/http"
	"public.sunibas.cn/go_utils_public/src/main/MainTools"
	"public.sunibas.cn/go_utils_public/src/utils/AboutServer"
	"public.sunibas.cn/go_utils_public/src/utils/DirAndFile"
	"strconv"
)

type FileOpEntity struct {
	Path string `json:"path"`
	//Type string `json:"type"` // file、dir
	Child bool `json:"child"`
	Zip bool `json:"zip"`
}
var FileOpActions []MainTools.ActionAtom = []MainTools.ActionAtom{
	{
		DearFn:      GetSubFiles,
		Method:      "GetSubFiles",
		Description: "获取文件",
		Params:      []string{"path:文件路径"},
	},
}

func GetSubFiles(writer http.ResponseWriter,s string,config MainTools.Config) {
	rObj := AboutServer.ReturnObj{}
	foe := FileOpEntity{}
	err := json.Unmarshal([]byte(s),&foe)
	if err != nil {
		rObj.SetFail("参数解析错误，err:" + err.Error()).Send(writer)
	} else {
		if t,e := DirAndFile.CheckTypeAndExist(foe.Path);e == nil && t == 1 {
			if foe.Child {
				if foe.Zip {
					fns := DirAndFile.GetAllSubDirOrFile(foe.Path)
					rObj.SetContent(fns).SetSuccess("").Send(writer)
				} else {
					fns := DirAndFile.GetAllSubDirOrFileNotZip(foe.Path)
					rObj.SetContent(fns).SetSuccess("").Send(writer)
				}
			} else {
				fns := DirAndFile.GetSubDirOrFile(foe.Path)
				rObj.SetContent(fns).SetSuccess("").Send(writer)
			}
		} else {
			rObj.SetFail("错误" + e.Error() + "，类型为（" + strconv.Itoa(t) + "）[-1是文件，0表示不存在]").Send(writer)
		}
	}
}
