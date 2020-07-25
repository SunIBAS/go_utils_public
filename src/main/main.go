package main

import (
	"os"
	"public.sunibas.cn/go_utils_public/src/main/InitHttp"
	_ "public.sunibas.cn/go_utils_public/statik"
)

func main() {
	if len(os.Args) > 2 {
		InitHttp.InitHttp(os.Args[1])
	} else {
		InitHttp.InitHttp("")
	}
}