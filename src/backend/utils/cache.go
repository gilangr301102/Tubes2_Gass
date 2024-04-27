package utils

import (
	"net/http"
	"time"

	lru "github.com/hashicorp/golang-lru"
)

const CacheSize = MAX_MEMO_SIZE

var client = &http.Client{
	Timeout: time.Second * TIMEOUT,
	Transport: &http.Transport{
		MaxIdleConnsPerHost: MAX_HOS_CONNECTION,
	},
}
var urlCache = make(map[string][]Node)
var sem = make(chan struct{}, SIZE_OF_CONTAINER)

var Articles *lru.Cache

func init() {
	cache, err := lru.New(CacheSize)
	PanicIfError(err)
	Articles = cache
}
