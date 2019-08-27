// A basic HTTP server.
// By default, it serves the current working directory on port 8080.
package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		return
	})
	http.Handle("/", http.FileServer(http.Dir(*dir)))
	err := http.ListenAndServe(*listen, nil)
	log.Fatalln(err)
}
