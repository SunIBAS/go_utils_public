package Console

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"syscall"
)

func RunCommand(name string,args []string) {
	cmd := exec.Command(name,args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Start(); err != nil {   // 运行命令
		fmt.Println(err)
	}
}

func RunBatFile(fileName string) {
	RunCommand("cmd",[]string { "/c","start", fileName })
}

func RunNodejs(filename string,nodepath string) (string,error) {
	cmd := exec.Command("cmd")
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in//绑定输入
	var out bytes.Buffer
	cmd.Stdout = &out //绑定输出
	cmd.Stderr = &out //绑定输出
	if len(nodepath) == 0 {
		nodepath = "node"
	}
	go func() {
		//in.WriteString("@echo off\n")
		//in.WriteString("cls\n")
		in.WriteString(strings.Join([]string{
			nodepath,
			"\"" + filename + "\"",
			"\n",
		}," "))//写入你的命令，可以有多行，"\n"表示回车
	}()
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cmd.Args)
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	return out.String(),err
}

// 在默认浏览器打开链接
func OpenUrl(url string)  {
	RunBatFile(url)
}
// 指定程序来打开文件
func OpenProgramWithFile(program,file string) {
	cmd := exec.Command("cmd","/C","start",program, file)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	err = cmd.Wait()
}
func OpenWordFile(filePath string) {
	OpenProgramWithFile("winword",filePath)
}
func OpenExcelFile(filePath string) {
	OpenProgramWithFile("excel",filePath)
}
func OpenPptFile(filePath string) {
	OpenProgramWithFile("powerpnt",filePath)
}
func OpenHtmlFileByChrome(filePath string) {
	OpenProgramWithFile("chrome",filePath)
}
func OpenTxtFileByNotepad(filePath string) {
	OpenProgramWithFile("notepad",filePath)
}
func OpenZipFileByRar(filePath string) {
	OpenProgramWithFile("winrar",filePath)
}