package routes

import (
	"fmt"
	"net/http"
)

func Bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, you hit bar!")

}
