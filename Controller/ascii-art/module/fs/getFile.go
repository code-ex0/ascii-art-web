package fs

import (
	"io/ioutil"
	"os"
	"strings"
)

func GetAlphabetFile(args []string) (result []string) {
	file := "standard.txt"
	if len(args) > 1 {
		switch args[1] {
		case "standard":
			file = "standard.txt"
		case "shadow":
			file = "shadow.txt"
		case "thinkertoy":
			file = "thinkertoy.txt"
		}
	}
	temp, _ := os.Open("Controller/ascii-art/file/" + file)
	temps, _ := ioutil.ReadAll(temp)
	temp.Close()
	return strings.Split(string(temps), "\r\n")
}
