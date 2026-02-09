package webcrawler

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type WebCrawler struct {
}

type stack struct {
	data map[string]bool
	mu   sync.Mutex
}

func (s *stack) len() int {
	return len(s.data)
}
func (s *stack) pop() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	for k, v := range s.data {
		if v == false {
			s.data[k] = true
			return k
		}
	}
	return ""
}
func (s *stack) push(url string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[url]; ok {
		return
	}
	s.data[url] = false
}
func (s *stack) getAll() map[string]bool {
	return s.data
}
func (w *WebCrawler) Start(seedURL string) {
	s := &stack{
		data: make(map[string]bool),
	}
	s.push(seedURL)
	w.process(seedURL, s)
}

func (w *WebCrawler) process(seedURL string, s *stack) {
	ch := make(chan string)
	var wg sync.WaitGroup
	workerCount := 5
	workerInProgress := 0
	var workerInProgressMu sync.Mutex
	wg.Add(workerCount)
	for range workerCount {
		go func() {
			defer wg.Done()
			for val := range ch {
				w.crawl(seedURL, val, s)
				workerInProgressMu.Lock()
				workerInProgress--
				workerInProgressMu.Unlock()
			}
		}()
	}

	for {
		url := s.pop()
		if url == "" && workerInProgress == 0 {
			fmt.Println("no more to process")
			break
		}
		if len(url) > 0 {
			fmt.Println("processing ", url)
			workerInProgressMu.Lock()
			workerInProgress++
			workerInProgressMu.Unlock()
			ch <- url
		}
	}

	close(ch)
	wg.Wait()
	//print/return all urls
	fmt.Println("urls found: ")
	for k := range s.getAll() {
		fmt.Println(k)
	}
}

func (w *WebCrawler) crawl(seed, url string, st *stack) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error creating request")
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making the cal")
		return
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("error reading res body")
		return
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		if link, ok := s.Attr("href"); ok && len(link) > 0 {
			if strings.HasPrefix(link, "https") && !strings.HasPrefix(link, seed) {
				return
			}
			if !strings.HasPrefix(link, "https") {
				link = strings.TrimSuffix(seed, "/") + "/" + strings.TrimPrefix(link, "/")
			}
			st.push(link)
		}
	})
}
