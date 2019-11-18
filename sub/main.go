package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/mcaci/calc/sub/serv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()

	srv := serv.NewSubService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := serv.Endpoints{
		SubEndpoint: serv.MakeSubEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("sub is listening on port:", *httpAddr)
		handler := serv.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
