package Controller

import (
	"./struct"
	"log"
	"net/http"
	"os"
)

var Config, _ = new(_struct.Config).NewConfig()

func Server() {
	logFile, err := os.OpenFile(Config.Server.Path.PathServer+Config.Server.Log.Path+Config.Server.Log.FileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	server := http.NewServeMux()
	server.HandleFunc(Config.Link.Home.UrlHome, Home)
	server.HandleFunc(Config.Link.Reverse.UrlReverse, Reverse)
	server.HandleFunc(Config.Link.Output.UrlOutput, Output)
	server.HandleFunc(Config.Link.Print.UrlPrint, PrintAscii)
	server.Handle(Config.Link.DownloadsFiles.UrlDownloadsFiles, http.StripPrefix(Config.Link.DownloadsFiles.UrlDownloadsFiles, http.FileServer(http.Dir(Config.Link.DownloadsFiles.PathDownloadsFiles))))
	server.Handle(Config.Link.LoadsAssets.UrlLoadsAssets, http.StripPrefix(Config.Link.LoadsAssets.UrlLoadsAssets, http.FileServer(http.Dir(Config.Link.LoadsAssets.PathLoadsAssets))))
	log.Fatal(http.ListenAndServe(Config.Server.Host+":"+Config.Server.Port, RequestLogger(server)))
}
