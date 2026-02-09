package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gorilla/mux"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	if err := mainErr(ctx, &wg); err != nil {
		log.Fatalf("fatal")
		return
	}

	signals := make(chan os.Signal, 2)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	cancel()
	wg.Wait()

}

func mainErr(ctx context.Context, wg *sync.WaitGroup) error {
	router := mux.NewRouter()
	router.HandleFunc("1", func(w http.ResponseWriter, r *http.Request) {})
	router.HandleFunc("2", func(w http.ResponseWriter, r *http.Request) {})
	router.HandleFunc("3", func(w http.ResponseWriter, r *http.Request) {})
	router.HandleFunc("4", func(w http.ResponseWriter, r *http.Request) {})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
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
