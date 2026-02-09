package shorturl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/esoptra/go-prac/shorturl/hashgenerator"
	"github.com/esoptra/go-prac/shorturl/store"
)

type URLShortener struct {
	store        store.Store
	urlGenerator hashgenerator.Generator
}

func NewURLShortener() *URLShortener {
	store := store.NewMapStore()
	return &URLShortener{
		store:        store,
		urlGenerator: hashgenerator.NewHashGenerator(8, 100, store),
	}
}

type URLShortRequest struct {
	URL string `json:"url"`
}

func (u *URLShortener) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/new":
		u.createShortURL(w, r)
	default:
		fmt.Println("url patht", r.URL.Path)
		val, err := u.store.RedirectStore().Get(strings.TrimPrefix(r.URL.Path, "/"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Println("redirec ", val)
		http.Redirect(w, r, val.(string), http.StatusPermanentRedirect)
	}
}

func (u *URLShortener) createShortURL(w http.ResponseWriter, r *http.Request) {
	req := &URLShortRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	hash := u.store.HashStore().GetHash()
	if hash == "" {
		fmt.Println("shouldn't happen")
	}
	u.store.RedirectStore().Create(hash, req.URL)
	w.Write([]byte(fmt.Sprintf(`{"url":"http://localhost:8000/%s"}`, hash)))

}
