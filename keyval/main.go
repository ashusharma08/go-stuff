package main

import (
	"fmt"
	"time"

	"github.com/esoptra/go-prac/keyval/keyval"
)

func main() {
	store := keyval.NewStore()
	fmt.Println("got store")
	fmt.Println(store.Get("Ashish"))
	store.Set("Ashish", 123)
	fmt.Println(store.Get("Ashish"))

	store.Set("Ashish", struct {
		Name     string
		LastName string
	}{
		Name:     "Ashish",
		LastName: "Sharma",
	})
	fmt.Println(store.Get("Ashish"))

	store.Del("Ashish")
	fmt.Println("null, ", store.Get("Ashish"))

	store.Set("AnotherKey", "AnotherValue", keyval.WithExpiry(2*time.Second))
	fmt.Println(store.Get("AnotherKey"))
	time.Sleep(3 * time.Second)
	fmt.Println(store.Get("AnotherKey"))
}
