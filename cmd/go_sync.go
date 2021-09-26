package main

import (
	"sync"

	"github.com/brunoshiroma/go-sync/internal"
)

func main() {
	var (
		wg         *sync.WaitGroup = &sync.WaitGroup{}
		errors     chan error
		httpServer internal.HttpServer
		udpServer  internal.UDPServer
	)
	errors = make(chan error, 100)
	httpServer = internal.NewHttpServer("0.0.0.0", 8080)
	udpServer = internal.NewUDPServer()
	wg.Add(1)
	go httpServer.Start(wg, errors)
	wg.Add(1)
	go udpServer.Start(wg, errors)
	wg.Wait()
	close(errors)
	printErrors(errors)
}

func printErrors(errors chan error) {
	for err := range errors {
		internal.LoggerS.Errorf("ERRORS: %#v", err)
	}
}
