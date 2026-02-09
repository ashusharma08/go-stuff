package hashgenerator

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/esoptra/go-prac/shorturl/store"
)

type HashGenerator struct {
	store      store.Store
	data       []string
	mu         sync.Mutex
	bufferSize int
	hashLength int
}

func NewHashGenerator(hashLength int, bufferSize int, store store.Store) *HashGenerator {
	h := &HashGenerator{
		data:       make([]string, bufferSize),
		bufferSize: bufferSize,
		hashLength: hashLength,
		store:      store,
	}
	go h.generateHashes()
	return h
}

func (h *HashGenerator) generateHashes() {
	fmt.Println("starting to generate hashses")
	for {
		if h.store.HashStore().GetSize() == h.bufferSize {
			time.Sleep(5 * time.Second)
		}

		randBytes := make([]byte, 8)
		_, err := rand.Read(randBytes)
		if err != nil {
			continue
		}
		hash := sha256.Sum256([]byte(randBytes))
		num := binary.BigEndian.Uint64(hash[:8])
		hs := base62encode(num)

		h.store.HashStore().SaveHash(hs)
	}
}

var base62items = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func base62encode(num uint64) string {
	if num == 0 {
		return string(base62items[0])
	}
	var sb strings.Builder
	for num > 0 {
		val := num % 62
		sb.WriteByte(base62items[val])
		num /= 62
	}
	result := sb.String()
	runes := []rune(result)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
