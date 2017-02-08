package main

import (
	"net/http"
	"time"
)

var routes []string

type customRouter struct {
}

func (customRouter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 1)
	rw.WriteHeader(http.StatusAccepted)
}

func main() {
	var cr customRouter
	server := &http.Server{
		Addr:           ":9000",
		Handler:        cr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
