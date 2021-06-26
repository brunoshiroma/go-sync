package main

import (
	"log"
	"net/http"
	"time"

	"github.com/brunoshiroma/go-sync/embed_html"
)

type indexHandler struct {
}

func (i *indexHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/" {
		response.Header().Add("Location", "/app/")
		response.WriteHeader(http.StatusMovedPermanently)
	}
}

func main() {
	i := &indexHandler{}
	http.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.FS(embed_html.Content))))
	http.Handle("/", i)
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
