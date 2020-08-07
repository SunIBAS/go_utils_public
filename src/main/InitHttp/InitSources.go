package InitHttp

import (
	"io/ioutil"
	"log"
	"net/http"
	"public.sunibas.cn/go_utils_public/src/utils/DirAndFile"
)

func InitSources(root string,statikFS http.FileSystem)  {
	files := []string {
		"Resources\\Nodejs\\formatNumber.js",
		"Resources\\Nodejs\\getDays.js",
	}
	for _,f := range files {
		writeToFile(root,f,statikFS)
	}
	DirAndFile.CreateFileBeforeCheckAndCreate(root + "tmp\\test.txt")
}

func writeToFile(root,filename string,statikFS http.FileSystem)  {
	fullPath := root + filename
	DirAndFile.CreateFileBeforeCheckAndCreate(fullPath)
	r, err := statikFS.Open("\\" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	DirAndFile.WriteWithIOUtil(fullPath,string(contents))
}
