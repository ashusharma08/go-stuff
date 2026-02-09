package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"slices"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/esoptra/go-prac/loadbalancer/utils"
	"github.com/sirupsen/logrus"
)

type Server struct {
	URL       string
	IsHealthy bool
	LastCheck time.Time
}

type LoadBalancer struct {
	servers []*Server
	mu      sync.Mutex
	current int
}

func NewLoadBalancer(urls []string) *LoadBalancer {
	l := &LoadBalancer{
		servers: make([]*Server, len(urls)),
		current: 0,
	}
	for idx, item := range urls {
		l.servers[idx] = &Server{
			URL: item,
		}
	}

	go l.monitorHealth()
	return l
}

type AddServerRequest struct {
	URL string `json:"url"`
}

func (l *LoadBalancer) AddNewServer(r *http.Request) error {
	if r.Method != http.MethodPut {
		return fmt.Errorf("Invalid method")
	}
	req := &AddServerRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || len(req.URL) == 0 {
		return err
	}
	_, err = url.Parse(req.URL)
	if err != nil {
		return err
	}
	l.mu.Lock()
	if slices.ContainsFunc(l.servers, func(s *Server) bool {
		return s.URL == req.URL
	}) {
		return nil
	}
	l.servers = append(l.servers, &Server{
		URL: req.URL,
	})
	l.mu.Unlock()
	return nil
}

func (l *LoadBalancer) monitorHealth() {
	client := utils.NewClient(utils.WithTimeout(2 * time.Second))
	for {
		l.mu.Lock()
		for _, s := range l.servers {
			_, err := client.Get(strings.TrimSuffix(s.URL, "/") + "/health-check")
			if err != nil {
				fmt.Println("health check error ", err)
				s.IsHealthy = false
			} else {
				s.IsHealthy = true
			}
			s.LastCheck = time.Now()
		}
		l.mu.Unlock()
		time.Sleep(5 * time.Second)
	}
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/add-server":
		if err := lb.AddNewServer(r); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	default:
		lb.serveRequest(w, r)
	}
}

func (lb *LoadBalancer) getNextServer() *Server {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	nextIdx := func() int {
		lb.current++
		if lb.current >= len(lb.servers) {
			lb.current = 0
		}
		return lb.current
	}
	for {
		idx := nextIdx()
		if lb.servers[idx].IsHealthy {
			return lb.servers[idx]
		}
	}
}

func (lb *LoadBalancer) serveRequest(w http.ResponseWriter, r *http.Request) {
	server := lb.getNextServer()
	if server == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	client := utils.NewClient()
	res, err := client.Do(server.URL, r)
	if err != nil {
		fmt.Println(" err ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	if res != nil {
		defer res.Body.Close()
		io.Copy(w, res.Body)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)

}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	if err := mainErr(ctx, &wg); err != nil {
		fmt.Println("fatal ", err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
	cancel()
	wg.Wait()
}

func mainErr(ctx context.Context, wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()
	lb := NewLoadBalancer([]string{"localhost:8080", "localhost:8081"})
	server := http.Server{
		Addr:    ":8000",
		Handler: lb,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				logrus.Infof("Server closed")
			} else {
				logrus.Errorf("something is wrong %#v", err)
			}
		}
	}()
	<-ctx.Done()
	server.Shutdown(ctx)
	return nil
}
