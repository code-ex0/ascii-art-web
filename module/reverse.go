package module

import (
	"./ascii-art"
	"./struct"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Reverse(w http.ResponseWriter, r *http.Request) {
	var sizeWindows int
	var reverse = new(_struct.DataReverse)
	tmpl := template.Must(template.ParseFiles(Config.Server.Path.PathServer + Config.Link.Reverse.PathReverseHtml))
	url := r.URL.Query()
	if url.Get("a") != "" && url.Get("f") != "" {
		if url.Get("a") == "use-file" {
			sizeWindows, _ = strconv.Atoi(url.Get("w"))
			reverse.FileName = url.Get("f")
			reverse.Result = ascii_art.Run_ascii_art([]string{"--reverse=" + reverse.FileName}, sizeWindows/8)
			reverse.Cat = ascii_art.Run_ascii_art([]string{reverse.Result, autoDetectType(reverse.FileName)}, sizeWindows/8)
		} else if url.Get("a") == "delete" {
			err := os.Remove(Config.Server.Path.PathServer + Config.Server.Path.PathOutputFile + url.Get("f"))
			if err != nil {
				log.Print(err)
			}
		} else if url.Get("a") == "download" {
		}

	}
	reverse.Files = getFile(Config.Server.Path.PathServer + Config.Server.Path.PathOutputFile)
	_ = tmpl.Execute(w, struct {
		Active      string
		ReverseData *_struct.DataReverse
		Url         *_struct.Config
	}{Active: "reverse", ReverseData: reverse, Url: Config})
	reverse.FileName = ""
	reverse.Cat = ""
	reverse.Result = ""
}

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
