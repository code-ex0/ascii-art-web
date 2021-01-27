package Controller

import (
	"./ascii-art"
	"./struct"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != Config.Link.Home.UrlHome {
		http.NotFound(w, r)
		return
	}
	tmpl := template.Must(template.ParseFiles(Config.Server.Path.PathServer + Config.Link.Home.PathHomeHtml))
	_ = tmpl.Execute(w, struct {
		Active string
		Url    *_struct.Config
	}{Active: "home", Url: Config})
}

func PrintAscii(w http.ResponseWriter, r *http.Request) {
	var sizeWindows int
	var printO = new(_struct.DataPrint)
	tmpl := template.Must(template.ParseFiles(Config.Server.Path.PathServer + Config.Link.Print.PathHomeHtml))
	if (r.FormValue("print-text") != "" && r.FormValue("select_type-text") != "" && r.FormValue("yon") != "" && r.FormValue("select-align") != "") && (r.FormValue("select-color") != "" ||
		r.FormValue("advanced-color") != "") {
		printO.Text = r.FormValue("print-text")
		printO.TextType = r.FormValue("select_type-text")
		printO.Color = r.FormValue("select-color")
		printO.ColorType = r.FormValue("yon")
		printO.Align = r.FormValue("select-align")
		sizeWindows, _ = strconv.Atoi(r.FormValue("width-window"))
		if printO.ColorType == "advanced" {
			printO.Color = r.FormValue("advanced-color")
		}
		printO.Result = ascii_art.Run_ascii_art([]string{printO.Text, printO.TextType, "--align=" + printO.Align}, sizeWindows/8)
	}
	_ = tmpl.Execute(w, struct {
		Active    string
		PrintData *_struct.DataPrint
		Url       *_struct.Config
	}{Active: "print", PrintData: printO, Url: Config})
	printO.Result = ""
}

func Output(w http.ResponseWriter, r *http.Request) {
	var sizeWindows int
	var output = new(_struct.DataOutput)
	tmpl := template.Must(template.ParseFiles(Config.Server.Path.PathServer + Config.Link.Output.PathOutputHtml))
	if r.FormValue("output-text") != "" && r.FormValue("output-file_name") != "" && r.FormValue("output-select_type-text") != "" {
		output.Text = r.FormValue("output-text")
		output.FileName = r.FormValue("output-file_name")
		output.TypeText = r.FormValue("output-select_type-text")
		sizeWindows, _ = strconv.Atoi(r.FormValue("width-window"))
		ascii_art.Run_ascii_art([]string{output.Text, output.TypeText, "--output=" + output.FileName}, 0)
		output.Result = ascii_art.Run_ascii_art([]string{output.Text, output.TypeText}, sizeWindows/8)
	}
	_ = tmpl.Execute(w, struct {
		Active     string
		OutputData *_struct.DataOutput
		Url        *_struct.Config
	}{Active: "output", OutputData: output, Url: Config})
	output.Result = ""
}

func Reverse(w http.ResponseWriter, r *http.Request) {
	var sizeWindows int
	var reverse = new(_struct.DataReverse)
	tmpl := template.Must(template.ParseFiles(Config.Server.Path.PathServer + Config.Link.Reverse.PathReverseHtml))
	if r.Method == "POST" {
		err := r.ParseMultipartForm(200000)
		if err != nil {
			return
		}
		formdata := r.MultipartForm
		files := formdata.File["multiplefiles"]
		for i, _ := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				return
			}
			out, err := os.Create(Config.Server.Path.PathServer + Config.Server.Path.PathOutputFile + files[i].Filename)
			defer out.Close()
			if err != nil {
				return
			}
			_, err = io.Copy(out, file)
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}

		}
	} else {
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
			}
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
