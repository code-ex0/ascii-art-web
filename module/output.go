package module

import (
	"./ascii-art"
	"./struct"
	"html/template"
	"net/http"
	"strconv"
)

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
