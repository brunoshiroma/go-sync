package main

import (
	"log"
	"net/http"
	"time"

	"github.com/brunoshiroma/go-sync/embed_html"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(embed_html.Content))))
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
