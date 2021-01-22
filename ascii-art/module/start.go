package module

import (
	"./align"
	"./fs"
	"./output"
	"./reverse"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var sizeWindows int

func Start(args []string, windows int) string {
	sizeWindows = windows
	if len(args) >= 1 {
		param := getParam(args)
		if v, found := param["align"]; found {
			return align.Align(PrintSentence(GetSentence(GetAlphabet(fs.GetAlphabetFile(args)), true, args[0])), v, sizeWindows)
		} else if v, found := param["output"]; found {
			if strings.HasSuffix(v, ".txt") {
				output.Output(v, GetSentence(GetAlphabet(fs.GetAlphabetFile(args)), false, args[0]))
			}
		} else if v, found := param["reverse"]; found {
			if strings.HasSuffix(v, ".txt") {
				f, err := os.Open("ascii-art/output/" + v)
				if err != nil {
					log.Fatal(err)
				}
				temp, _ := ioutil.ReadAll(f)
				f.Close()
				return reverse.Reverse(strings.Split(string(temp), "\n"), GetAlphabet(autoDetectTypeFile(v)))
			}
		} else {
			return align.Align(PrintSentence(GetSentence(GetAlphabet(fs.GetAlphabetFile(args)), true, args[0])), "left", sizeWindows)
		}
	}
	return ""
}

func getParam(args []string) map[string]string {
	param := make(map[string]string)
	for _, i := range args {
		if strings.HasPrefix(i, "--") {
			if strings.Contains(i, "output") {
				param["output"] = i[strings.Index(i, "=")+1:]
			} else if strings.Contains(i, "reverse") {
				param["reverse"] = i[strings.Index(i, "=")+1:]
			} else if strings.Contains(i, "align") {
				param["align"] = i[strings.Index(i, "=")+1:]
			}
		}
	}
	return param
}

func autoDetectTypeFile(file string) []string {
	temp, _ := os.Open("ascii-art/output/" + file)
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
	var typefile string
	if standard {
		typefile = "standard.txt"
	} else if shadow {
		typefile = "shadow.txt"
	} else if thinkertoy {
		typefile = "thinkertoy.txt"
	}

	a, _ := os.Open("ascii-art/file/" + typefile)
	b, _ := ioutil.ReadAll(a)
	a.Close()
	return strings.Split(string(b), "\r\n")
}
