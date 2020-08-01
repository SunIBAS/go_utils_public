package InitHttp

import (
	"fmt"
	"github.com/rakyll/statik/fs"
	"github.com/semicircle/gocors"
	"log"
	"net/http"
	"public.sunibas.cn/go_utils_public/src/main/MainTools"
	"time"
)

var (
	config = MainTools.Config{}
)

func InitHttp(configPath string) {
	config = Entrance(configPath)
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	cors := gocors.New()
	http.Handle("/public/", cors.Handler(http.StripPrefix("/public/", http.FileServer(statikFS))))
	// 默认访问主页
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/public/views/index.html", http.StatusFound)
	})
	InitAction()
	InitSocket()
	//
	InitPubWeb()
	srv := &http.Server{
		Addr:           config.Port,
		Handler:        nil,
		ReadTimeout:    time.Duration(5) * time.Minute,
		WriteTimeout:   time.Duration(5) * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
