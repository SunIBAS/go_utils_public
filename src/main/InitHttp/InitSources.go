package InitHttp

import (
	"encoding/json"
	"public.sunibas.cn/go_utils_public/src/Datas"
	"public.sunibas.cn/go_utils_public/src/main/Sqls"
	"public.sunibas.cn/go_utils_public/src/utils/DirAndFile"
)

type fileRec struct {
	FileName string `json:"FileName"`
	Content string `json:"Content"`
}
func InitSourcesFromDB() {
	records := Sqls.ParseRowBySelectRecords(config.DB," where `tag`='files'",config.Tables["records"])
	var fr fileRec
	// config.CwdPath
	for _,rec := range records {
		err := json.Unmarshal([]byte(Datas.DecodeBase64AndURI(rec.Content)),&fr)
		filename := config.CwdPath + DirAndFile.Filepath_Separator + fr.FileName
		if err == nil {
			if t,e := DirAndFile.CheckTypeAndExist(filename);
				e == nil {
					if t == 0 {
						// 只有这种情况时创建文件
						DirAndFile.WriteWithIOUtil(filename,fr.Content)
					}
			}
		}

	}
}