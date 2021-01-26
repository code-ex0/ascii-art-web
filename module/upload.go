package module

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

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
	fmt.Println(Config.Server.Host + ":" + Config.Server.Port + Config.Link.Reverse.UrlReverse)
	http.Redirect(w, r, Config.Server.Host+":"+Config.Server.Port+Config.Link.Reverse.UrlReverse, http.StatusSeeOther)

}
