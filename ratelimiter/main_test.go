package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"
)

func Test_ratelimiter(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(10)
	i := 0
	for range 10 {
		go func(in int) {
			defer wg.Done()
			req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8000", nil)
			if err != nil {
				t.Fatal("error should not be nil")
			}
			req.Header.Set("userid", fmt.Sprintf("user%d", in))
			for range 11 {
				res, err := http.DefaultClient.Do(req)
				if err != nil {
					t.Fatal("error in making do")
				}
				if res.StatusCode != http.StatusOK {
					t.Fatalf("error in status code. expected 200 got %d", res.StatusCode)
				}
				resp, err := io.ReadAll(res.Body)
				if err != nil {
					t.Fatalf("rsponse error %#v ", err)
				}
				fmt.Println("___________", string(resp))
			}
		}(i)
		i++
	}
	wg.Wait()
}
