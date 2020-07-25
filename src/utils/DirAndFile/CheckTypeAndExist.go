package DirAndFile

import "os"

const (
	File = -1
	NoExist = 0
	Dir = 1
)

/**
 * 如果文件存在则返回 -1 表示文件，1 表示文件夹
 * 文件不存在则返回 0
 */
func CheckTypeAndExist(file string) (int,error)  {
	if len(file) == 0 {
		return 0,nil
	}
	s, err := os.Stat(file)
	if err == nil {
		if s.IsDir() {
			return 1,nil
		} else {
			return -1, nil
		}
	}
	if os.IsNotExist(err) {
		return 0, nil
	}
	return 0, err
}
