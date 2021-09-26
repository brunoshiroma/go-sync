package internal

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/brunoshiroma/go-sync/embed_html"
)

type HttpServer interface {
	Start(wg *sync.WaitGroup, errors chan error)
	Stop() error
}

type SimpleHttpServer struct {
	host       string
	port       int
	serverMux  *http.ServeMux
	httpServer *http.Server
}

func (h *SimpleHttpServer) Start(wg *sync.WaitGroup, errors chan error) {
	LoggerS.Infof("Starting http server on %s:%d", h.host, h.port)
	errors <- h.httpServer.ListenAndServe()
	wg.Done()
}

func (h *SimpleHttpServer) Stop() (err error) {
	LoggerS.Info("Stopping http server")
	err = h.httpServer.Shutdown(context.TODO())
	return
}

func (s *SimpleHttpServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	LoggerS.Debugw("RECEIVED REQUEST",
		"url", request.URL,
		"remoteAddress", request.RemoteAddr,
		"UA", request.Header.Get("User-Agent"))
	if request.URL.Path == "/" {
		response.Header().Add("Location", "/app/")
		response.WriteHeader(http.StatusMovedPermanently)
	}
}

func (h *SimpleHttpServer) setupMux() {
	h.serverMux = http.NewServeMux()
	h.serverMux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.FS(embed_html.Content))))
	h.serverMux.Handle("/", h)
}

func NewHttpServer(host string, port int) HttpServer {
	simpleHttpServer := &SimpleHttpServer{
		port: port,
		host: host,
	}

	simpleHttpServer.setupMux()

	simpleHttpServer.httpServer = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", host, port),
		Handler:        simpleHttpServer.serverMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return simpleHttpServer
}
