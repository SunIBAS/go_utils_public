package DirAndFile

import (
	"os"
	"os/exec"
	"path/filepath"
)

func GetCwd() string {
	file, _ := exec.LookPath(os.Args[0])
	fullPath, _ := filepath.Abs(file)
	paths, _ := filepath.Split(fullPath)
	return paths
}
