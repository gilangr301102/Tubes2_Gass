package utils

import (
	"net/http"
	"time"

	lru "github.com/hashicorp/golang-lru"
)

const CacheSize = 200000

var (
	httpClient = &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
		},
	}
	urlCache = make(map[string][]Node)
	sem      = make(chan struct{}, 500)
)

var Articles *lru.Cache

func init() {
	cache, err := lru.New(CacheSize)
	PanicIfError(err)
	Articles = cache
}
