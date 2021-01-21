package _struct

import "os"

type DataReverse struct {
	FileName string
	Files    []os.FileInfo
	Cat      string
	Result   string
}
