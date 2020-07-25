package Console

import (
	"fmt"
	"os/exec"
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