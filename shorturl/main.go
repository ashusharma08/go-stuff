package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/esoptra/go-prac/shorturl/shorturl"
)

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	err := mainErr(ctx, &wg)
	if err != nil {
		log.Fatalf("fatal err %#v", err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
	cancel()
	wg.Wait()
}

func mainErr(ctx context.Context, wg *sync.WaitGroup) error {
	server := http.Server{
		Addr:    ":8000",
		Handler: shorturl.NewURLShortener(),
	}
	wg.Add(1)
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
