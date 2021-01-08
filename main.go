package main

import (
	"./ascii-art"
	"fmt"
	"html/template"
	"net/http"
)

type dataPrint struct {
	Text      string
	TextType  string
	Color     string
	ColorType string
	Align     string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/index.html"))
		_ = tmpl.Execute(w, struct{ Active string }{Active: "home"})
	})

	http.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/index.html"))
		_ = tmpl.Execute(w, struct{ Active string }{Active: "reverse"})
	})
	http.HandleFunc("/output", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/index.html"))
		_ = tmpl.Execute(w, struct{ Active string }{Active: "output"})
	})
	http.HandleFunc("/print", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/index.html"))

		data := dataPrint{
			Text:      r.FormValue("print-text"),
			TextType:  r.FormValue("select_type-text"),
			Color:     r.FormValue("select-color"),
			ColorType: r.FormValue("yon"),
		}
		if data.ColorType == "advanced" {
			data.Color = r.FormValue("advanced-color")
		}
		fmt.Println(data)
		if data.Text != "" {
			args := []string{data.Text, data.TextType}
			_ = tmpl.Execute(w, struct {
				Active string
				Print  bool
				Text   string
				Data   dataPrint
			}{Active: "print", Print: true, Text: ascii_art.Run_ascii_art(args), Data: data})
		} else {
			_ = tmpl.Execute(w, struct {
				Active string
				Print  bool
				Data   dataPrint
			}{Active: "print", Print: false, Data: data})
		}
	})
	_ = http.ListenAndServe(":80", nil)
}
