package routes

import (
	"fmt"
	"net/http"
)

func Foo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Hello, you hit foo with GET!")
}
