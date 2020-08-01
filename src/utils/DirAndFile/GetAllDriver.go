package DirAndFile

import "os"

func GetDrives() (r []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ"{
		_, err := os.Open(string(drive)+":\\")
		if err == nil {
			r = append(r, string(drive))
		}
	}
	return
}