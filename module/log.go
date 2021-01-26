package module

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := NewLoggingResponseWriter(w)
		start := time.Now()
		targetMux.ServeHTTP(lrw, r)

		logs := fmt.Sprintf("%s\t\t%d\t\t%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			lrw.statusCode,
			http.StatusText(lrw.statusCode),
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
		log.Printf(logs)
		fmt.Println(logs)
	})
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
}
