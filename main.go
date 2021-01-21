package main

import (
	"./ascii-art"
	"./struct"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var reverse = new(_struct.DataReverse)
var printO = new(_struct.DataPrint)
var output = new(_struct.DataOutput)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	r.HandleFunc("/upload-file-reverse", uploadFile)
	r.HandleFunc("/reverse/{action}/{file}", getReverse)
	r.HandleFunc("/reverse", Reverse)
	r.HandleFunc("/output", Output)
	r.HandleFunc("/print", PrintAscii)
	_ = http.ListenAndServe(":80", r)
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/index.html"))
	_ = tmpl.Execute(w, struct{ Active string }{Active: "home"})
}

func Reverse(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/index.html"))
	if reverse.FileName != "" {
		reverse.Result = ascii_art.Run_ascii_art([]string{"--reverse=" + reverse.FileName})
		reverse.Cat = ascii_art.Run_ascii_art([]string{reverse.Result, autoDetectType(reverse.FileName)})
		fmt.Println(reverse)
	}
	reverse.Files = getFile("ascii-art/output")
	_ = tmpl.Execute(w, struct {
		Active      string
		ReverseData *_struct.DataReverse
	}{Active: "reverse", ReverseData: reverse})
	reverse.FileName = ""
	reverse.Cat = ""
	reverse.Result = ""
}

func Output(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/index.html"))
	if r.FormValue("output-text") != "" && r.FormValue("output-file_name") != "" && r.FormValue("output-select_type-text") != "" {
		output.Text = r.FormValue("output-text")
		output.FileName = r.FormValue("output-file_name")
		output.TypeText = r.FormValue("output-select_type-text")
		ascii_art.Run_ascii_art([]string{output.Text, output.TypeText, "--output=" + output.FileName})
		output.Result = ascii_art.Run_ascii_art([]string{output.Text, output.TypeText})
		fmt.Println(output)
	}
	_ = tmpl.Execute(w, struct {
		Active     string
		OutputData *_struct.DataOutput
	}{Active: "output", OutputData: output})
	output.Result = ""
}

func PrintAscii(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/index.html"))
	if (r.FormValue("print-text") != "" && r.FormValue("select_type-text") != "" && r.FormValue("yon") != "") && (r.FormValue("select-color") != "" ||
		r.FormValue("advanced-color") != "") {
		printO.Text = r.FormValue("print-text")
		printO.TextType = r.FormValue("select_type-text")
		printO.Color = r.FormValue("select-color")
		printO.ColorType = r.FormValue("yon")
		if printO.ColorType == "advanced" {
			printO.Color = r.FormValue("advanced-color")
		}
		printO.Result = ascii_art.Run_ascii_art([]string{printO.Text, printO.TextType})
		fmt.Println(printO)
	}
	_ = tmpl.Execute(w, struct {
		Active    string
		PrintData *_struct.DataPrint
	}{Active: "print", PrintData: printO})
	printO.Result = ""

}

func getFile(directory string) []os.FileInfo {
	files := make(map[string]string)
	file, err := ioutil.ReadDir(directory + "/")
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

func getReverse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["action"] == "delete" {
		err := os.Remove("ascii-art/output/" + vars["file"])
		if err != nil {
			fmt.Println(err)
		}
	} else if vars["action"] == "use-file" {
		reverse.FileName = vars["file"]
	}
	http.Redirect(w, r, "/reverse", http.StatusSeeOther)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
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
		out, err := os.Create("ascii-art/output/" + files[i].Filename)
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
	http.Redirect(w, r, "/reverse", http.StatusSeeOther)
}

func autoDetectType(file string) (result string) {
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
	if standard {
		result = "standard"
	} else if shadow {
		result = "shadow"
	} else if thinkertoy {
		result = "thinkertoy"
	}
	fmt.Println(result)
	return
}
