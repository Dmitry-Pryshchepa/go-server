package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"./routes"
	"./templates"
)

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

func templateHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}


func setRoutes(h *http.ServeMux) {

	h.HandleFunc("/foo", routes.Foo)

	h.HandleFunc("/bar", routes.Bar)

	h.Handle("/main", &templates.TemplateHandler{Filename: "main.html"})

	h.Handle("/app", &templates.TemplateHandler{Filename: "app.html"})

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "You're lost, go home")
	})
}

func main() {

	h := http.NewServeMux()
	setRoutes(h)
	l:= logger(h)

	server := &http.Server{
		Addr:           ":9000",
		Handler:        l,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}

