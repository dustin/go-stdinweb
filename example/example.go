package main

import (
	"fmt"
	"github.com/dustin/go-stdinweb"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %v, you wanted %q\n",
			r.RemoteAddr, r.URL.Path)
	})
	stdinweb.ServeStdin(http.Server{})
}
