package main

import (
	"fmt"
	"github.com/dustin/go-stdinweb"
	"net/http"
)

func main() {
	s := http.Server{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q\n", r.URL.Path)
	})
	stdinweb.ServeStdin(s)
}
