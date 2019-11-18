package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/mcaci/mu-calc/div/serv"
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

	srv := serv.NewDivService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := serv.Endpoints{
		DivEndpoint: serv.MakeDivEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("div is listening on port:", *httpAddr)
		handler := serv.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
