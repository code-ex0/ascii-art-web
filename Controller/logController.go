package Controller

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}

func LogHTTP(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := statusWriter{ResponseWriter: w}
		handler.ServeHTTP(&sw, r)
		duration := time.Now().Sub(start)
		logs := fmt.Sprintf("%s\t\t%s\t\t%s\t\t%s\t\t%s\t\t%d\t\t%d\t\t%s\t\t%v",
			r.Host,
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			r.Proto,
			sw.status,
			sw.length,
			r.Header.Get("User-Agent"),
			duration)
		log.Printf(logs)
		fmt.Println(logs)
	}
}
