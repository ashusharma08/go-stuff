package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/esoptra/go-prac/restnethttp/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	if err := mainErr(ctx, &wg); err != nil {
		log.Fatalf("error starting server %#v", err)
		return
	}

	signals := make(chan os.Signal, 2)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	cancel()
	wg.Wait()
}

func mainErr(ctx context.Context, wg *sync.WaitGroup) error {
	wg.Add(1)
	server := http.Server{
		Addr:    ":8080",
		Handler: server.NewServer(),
	}

	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil {
			return
		}
	}()
	<-ctx.Done()
	server.Shutdown(ctx)
	return nil
}
