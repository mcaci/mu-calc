package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/mcaci/mu-calc/mul/serv"
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

	srv := serv.NewMulService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := serv.Endpoints{
		MulEndpoint: serv.MakeMulEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("mul is listening on port:", *httpAddr)
		handler := serv.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
