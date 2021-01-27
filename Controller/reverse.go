package Controller

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFile(directory string) []os.FileInfo {
	files := make(map[string]string)
	file, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	num := 1
	for _, f := range file {
		if strings.HasSuffix(f.Name(), ".txt") {
			files[strconv.Itoa(num)] = f.Name()
			num++
		}
	}
	return file
}

func autoDetectType(file string) (result string) {
	temp, _ := os.Open(Config.Server.Path.PathServer + Config.Server.Path.PathOutputFile + file)
	standard := false
	shadow := false
	thinkertoy := false
	temps, _ := ioutil.ReadAll(temp)
	temp.Close()
	for i := 0; i < len(temps)-1; i++ {
		if temps[i] == 'o' || temps[i] == '-' {
			shadow = false
			standard = false
			thinkertoy = true
		} else if temps[i] == ',' || temps[i] == ')' || temps[i] == '(' || temps[i] == 'V' || temps[i] == '/' || temps[i] == '\\' || temps[i] == '<' || temps[i] == '>' {
			shadow = false
			thinkertoy = false
			standard = true
		} else if temps[i] == '_' || temps[i+1] == '|' {
			thinkertoy = false
			shadow = true
		}
	}
	if standard {
		result = "standard"
	} else if shadow {
		result = "shadow"
	} else if thinkertoy {
		result = "thinkertoy"
	}
	return
}
