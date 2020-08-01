package DirAndFile

import (
	"os"
	"strings"
)

// p1 相对于 p2
// p1 = "c:\\first\\a\\b"
// p2 = "c:\\first\\c\\d"
// return = "./../../c/d"
func GetRelativePath(p1,p2 string) string {
	pp1 := strings.Split(p1,Filepath_Separator)
	pp2 := strings.Split(p2,Filepath_Separator)
	if pp1[0] == pp2[0] {
		str := []string{}
		// pp1 = [a,b,c,d,e]
		// pp2 = [a,b,f,g,h]
		// => [<-c,<-d,<-e,f,g,h]
		ind := 1
		for pp1[ind] == pp2[ind] {
			ind++
		}
		ind++
		for pp1Ind := ind;pp1Ind < len(pp1);pp1Ind++ {
			str = append(str,"..")
		}
		for pp2Ind := ind;pp2Ind < len(pp2);pp2Ind++ {
			str = append(str,pp2[pp2Ind])
		}
		return strings.Replace("./" + strings.Join(str,Filepath_Separator),"\\","/",-1)
	} else {
		return ""
	}
}

// 和当前目录的相对位置
func GetRelativePathWithCwd(p string) string {
	return GetRelativePath(GetCwd(),p)
}
// 和当前目录的相对位置
func GetRelativePathWithOSCwd(p string) string {
	dir,_ := os.Getwd()
	return GetRelativePath(dir,p)
}