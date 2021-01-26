package module

import (
	"./struct"
	"html/template"
	"net/http"
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
