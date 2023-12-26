package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	d := flag.String(
		"d",
		".",
		"directory from where files are served")
	p := flag.String(
		"p",
		"8080",
		"Port for serving files.")

	flag.Parse()

	http.Handle(
		"/",
		http.FileServer(
			http.Dir(*d)))

	log.Printf("Serving files from %s on port: %s\n", *d, *p)
	log.Fatal(http.ListenAndServe(":"+*p, nil))
}
