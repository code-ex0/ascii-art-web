package module

import (
	"./ascii-art"
	"./struct"
	"html/template"
	"net/http"
	"strconv"
)

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
